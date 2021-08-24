package cache

import "time"

type sweeper struct {
	isSweeping bool
	// start or stop sweep signal. If send true start sweep, false stop sweep.
	expireSignal chan bool
	// sweep interval for expire key.Default value is 5s.
	expireDuration time.Duration
	expireMap      map[string]time.Time
	// callback when expire key deleted.
	onExpireDelete func(key string)
}

func newSweeper(expireDuration time.Duration, onExpireDelete func(key string)) *sweeper {
	s := &sweeper{
		isSweeping:     false,
		expireMap:      make(map[string]time.Time),
		expireSignal:   make(chan bool),
		expireDuration: expireDuration,
		onExpireDelete: onExpireDelete,
	}
	if s.expireDuration == 0 {
		s.expireDuration = time.Second * 5
	}
	// create a goroutine to run sweep.
	go s.sweep()
	return s
}

func (s *sweeper) addExpireKey(key string, duration time.Duration) {
	s.tryStartSweep()
	s.expireMap[key] = time.Now().Add(duration)
}

func (s *sweeper) tryStartSweep() {
	if !s.isSweeping {
		s.isSweeping = true
		s.expireSignal <- true
	}
}

func (s *sweeper) tryStopSweep() {
	if s.isSweeping && len(s.expireMap) <= 0 {
		s.isSweeping = false
		s.expireSignal <- false
	}
}

func (s *sweeper) sweep() {
Restart:
	//blocking until signal comes
	select {
	case sig := <-s.expireSignal:
		if !sig {
			goto Restart
		}
	}
	//the start signal is coming.
	for {
		select {
		case sig := <-s.expireSignal:
			//if get false signal, need stop sweep.
			if !sig {
				goto Restart
			}
		default:
			for key, expireTime := range s.expireMap {
				// if expire time is before now, need delete this key and value in the cache.
				if expireTime.Before(time.Now()) {
					s.delete(key)
				}
			}
			time.Sleep(s.expireDuration)
			s.tryStopSweep()
		}
	}
}

func (s *sweeper) delete(key string) {
	if _, ok := s.expireMap[key]; ok {
		delete(s.expireMap, key)
		s.onExpireDelete(key)
	}
}

func (s *sweeper) isExpireKey(key string) bool {
	if _, ok := s.expireMap[key]; ok {
		return true
	}
	return false
}

func (s *sweeper) isExpired(key string) bool {
	if expireTime, ok := s.expireMap[key]; ok {
		if expireTime.Before(time.Now()) {
			return true
		}
	}
	return false
}
