package control

import (
	"sync"

	"github.com/go-vgo/robotgo"
)

type Media struct {
	mu        *sync.Mutex
	isPlaying bool
	isPaused  bool
}

func NewMedia() *Media {
	return &Media{
		mu:        &sync.Mutex{},
		isPlaying: false,
		isPaused:  false,
	}
}

func (m *Media) Play() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.isPlaying {
		return nil
	}
	if err := robotgo.KeyTap(robotgo.AudioPlay); err != nil {
		return err
	}
	m.isPlaying = true
	m.isPaused = false
	return nil
}

func (m *Media) Pause() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.isPaused {
		return nil
	}
	if err := robotgo.KeyTap(robotgo.AudioPause); err != nil {
		return err
	}
	m.isPlaying = false
	m.isPaused = true
	return nil
}

func (m *Media) Stop() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.isPaused {
		return nil
	}
	if err := robotgo.KeyTap(robotgo.AudioStop); err != nil {
		return err
	}
	m.isPlaying = false
	m.isPaused = true
	return nil
}

func (m *Media) Next() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := robotgo.KeyTap(robotgo.AudioNext); err != nil {
		return err
	}
	m.isPlaying = true
	m.isPaused = false
	return nil
}

func (m *Media) Previous() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := robotgo.KeyTap(robotgo.AudioPrev); err != nil {
		return err
	}
	m.isPlaying = true
	m.isPaused = false
	return nil
}
