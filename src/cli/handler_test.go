package cli

import (
	"testing"
)

func TestArgsHandler_SetKeyFromCli(t *testing.T) {

	arguments := []string{"passStore search -P -bedrin"}
	wantKeys := []uint8{80, 98, 101, 100, 114, 105, 110}
	//pf := ArgsHandler{arguments, wantKeys, []string{}}
	pf := ArgsHandler{arguments: arguments}
	pf.SetKeyFromCli()
	t.Log(pf.keys)
	if len(pf.keys) == 0 {
		t.Fatal( "Empty keys")
	}

	for i,value := range pf.keys {
			if wantKeys[i] != value {
				t.Fatalf( "%c %c\n",wantKeys[i], value)
			}
	}

}