package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"time"

	"github.com/H3Cki/wscsrv/control"
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media"
)

func connect() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8000/signal", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = conn.WriteJSON(map[string]any{
		"convid": "ci1",
		"type":   "create_answerer",
		"data": map[string]any{
			"name":          "rad1",
			"description":   "xd",
			"accesskey":     "ac1",
			"managementkey": "mk1",
		}})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("registered")

	for {
		resp := map[string]any{}
		if err := conn.ReadJSON(&resp); err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("\n\n[%s] : %s\n\n", resp["type"], resp)

		if resp["type"] == "agreement_offer" {
			sdp := resp["data"].(map[string]any)["sdp"].(string)
			agreementid := resp["data"].(map[string]any)["agreementid"].(string)
			offerername := resp["data"].(map[string]any)["offerername"].(string)

			answerSDP, err := connectionForOffer(sdp)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = conn.WriteJSON(map[string]any{
				"convid": "ci2",
				"type":   "accept_agreement",
				"data": map[string]any{
					"agreementid": agreementid,
					"sdp":         answerSDP,
				}})
			if err != nil {
				panic(err)
			}

			fmt.Printf("\naccepted agreement %s from %s\n", agreementid, offerername)
		}
	}

}

func connectionForOffer(offerSDP string) (string, error) {
	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if cErr := peerConnection.Close(); cErr != nil {
	// 		fmt.Printf("cannot close peerConnection: %v\n", cErr)
	// 	}
	// }()

	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		fmt.Printf("Peer Connection State has changed: %s\n", s.String())
	})

	peerConnection.OnDataChannel(func(c *webrtc.DataChannel) {
		c.OnOpen(func() {
			fmt.Println("Data Channel OPEN")
		})
		c.OnClose(func() {
			fmt.Println("Data Channel CLOSE")
		})
		c.OnError(func(err error) {
			fmt.Printf("\n\nerr: %s\n\n", err)
		})
		c.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Println(msg)
		})
	})

	peerConnection.OnTrack(func(tr *webrtc.TrackRemote, r *webrtc.RTPReceiver) {
		fmt.Println("streaming video")
		imagesC := make(chan image.Image, 1)
		scrn := &control.Screen{}

		// Initialize video stream settings
		videoTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: "image/jpeg"}, "video", "pion")
		if err != nil {
			// Handle error
			return
		}

		go func() {
			for {
				img, err := scrn.Screenshot()
				if err != nil {
					panic(err)
				}

				imagesC <- img
			}
		}()

		for img := range imagesC {
			var buf bytes.Buffer
			err = jpeg.Encode(&buf, img, nil)
			if err != nil {
				// Handle error
				continue
			}

			// Send image data over WebRTC
			err = videoTrack.WriteSample(media.Sample{Data: buf.Bytes(), Duration: time.Second / 30})
			if err != nil {
				// Handle error
				continue
			}

		}
	})

	if err := peerConnection.SetRemoteDescription(webrtc.SessionDescription{
		Type: webrtc.SDPTypeOffer,
		SDP:  offerSDP,
	}); err != nil {
		return "", err
	}

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		return "", err
	}

	if err := peerConnection.SetLocalDescription(answer); err != nil {
		return "", err
	}

	<-webrtc.GatheringCompletePromise(peerConnection)

	return peerConnection.LocalDescription().SDP, nil
}
