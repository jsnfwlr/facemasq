package logging

import (
	"testing"
)

func TestLogging(t *testing.T) {
	mapt := make(map[string]string)
	mapt["1"] = "one"
	mapt["two"] = "2"
	message := prepareMessage("%+v", mapt)
	t.Log(message)

	message = prepareMessage("%+[2]v - %[1]s", "extra", mapt)
	t.Log(message)

	message = prepareMessage("%.2f", 19.3246)
	t.Log(message)

	message = prepareMessage("%.2[1]f", 19.3246)
	t.Log(message)

	message = prepareMessage(1, 2, 3, 4)
	t.Log(message)

	logit := New("", "")
	var portList []int64
	for i := int64(1); i <= 1024; i++ {
		portList = append(portList, i)
	}

	for a := 0; a < 3; a++ {
		j := 0
		for i := range portList {
			if i != 0 && i%(len(portList)/10) == 0 {
				message = prepareMessage("%d\n", j)
				t.Log(message)
				logit.Debug1("! %d\n", j)
				j++

			}
		}

	}
	t.Error("Fake error to see stdOut")
}
