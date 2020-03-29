package main

import (
    "testing"
)

type test struct {
        in string
        out string
        proc procFunc
        shouldErr bool
        name string
}

func TestValidate(t *testing.T) {

    tests := []test {
        {"bad", "", nil, true, "filename with no suffix"},
        {"bad_ext.foo", "", nil, true, "filename with bad extension"},
        {"good_png.png", "good_png.dei", encode, false, "Good png name"},
        {"good_dei.dei", "good_dei.png", decode, false, "Good dei name"},
        {"caps_png.PNG", "caps_png.dei", encode, false, "Caps png name"},
        {"caps_dei.DEI", "caps_dei.png", decode, false, "Caps dei name"},
        {"mixed_png.PnG", "mixed_png.dei", encode, false, "Mixed-case png name"},
        {"mixed_dei.Dei", "mixed_dei.png", decode, false, "Mixed-case dei name"},
    }

    for _, tst := range tests {
        out, proc, err := validateFile(tst.in)
        if tst.shouldErr && err == nil {
            t.Errorf("Test <" + tst.name + "> did not fail when it should have")
        }
        if !tst.shouldErr && err != nil {
            t.Errorf("Test <" + tst.name + "> failed when it should have passed")
        }
        if tst.shouldErr {
            if out != "" {
                t.Errorf("Error result returned non-empty output string: " + out)
            }
            if proc != nil {
                t.Errorf("Error result returned non-nil process function")
            }
        } else {
            if out != tst.out {
                t.Errorf("Bad output string. Wanted <" + tst.out + "> got <"+ out + ">")
            }
            if proc == nil {
                t.Errorf("Proc func wanted one, got nil")
            }
        }
    }    
}