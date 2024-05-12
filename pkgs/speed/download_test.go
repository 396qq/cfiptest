package speed

import "testing"

func TestGetDownloadSpeed(t *testing.T) {
	st := &CFSpeedTest{
		EnableTLS:        true,
		SpeedTestTimeout: 5,
		MinSpeed:         4,
	}
	st.PreSetArgs()

	ip := "172.64.155.2"
	port := 443

	speed, _, err := st.getDownloadSpeed(ip, port)

	if err != nil {
		t.Errorf("getDownloadSpeed(%s, %d) returned an error: %v", ip, port, err)
	}

	expectedSpeed := 10.0 // You should replace this with the expected speed
	if speed != expectedSpeed {
		t.Errorf("getDownloadSpeed(%s, %d) = %f; want %f", ip, port, speed, expectedSpeed)
	}
}
