package main

import (
	"fmt"
	"testing"
)

func cmpBeam(lhsx int, lhsy int, lhsd beamDirection, rhsx int, rhsy int, rhsd beamDirection) error {
	if lhsx == rhsx && lhsy == rhsy && lhsd == rhsd {
		return nil
	}
	return fmt.Errorf("beams not equal: x=[%d, %d], y=[%d, %d], dir=[%d, %d]", lhsx, rhsx, lhsy, rhsy, lhsd, rhsd)
}

func compare(lhs beams, rhs beams) error {
	if lhs.len() != rhs.len() {
		return fmt.Errorf("beam sets are not the same length")
	}

	var err error
	for i := range lhs.len() {
		lhsx, lhsy, lhsd := lhs.get(i)
		rhsx, rhsy, rhsd := rhs.get(i)
		if cmp_err := cmpBeam(lhsx, lhsy, lhsd, rhsx, rhsy, rhsd); cmp_err != nil {
			err = fmt.Errorf("%w, %w", cmp_err, err)
		}
	}
	return err
}

func TestStep(t *testing.T) {
	bms := beams{}
	bms.append(0, 0, Left)
	bms.append(0, 0, Down)
	bms.append(1, 2, Down)
	bms.append(1, 2, Up)
	bms.append(1, 1, Right)
	output := stepBeamPosition(bms, 3, 3)
	expected := beams{}
	expected.append(0, 1, Down)
	expected.append(1, 1, Up)
	expected.append(2, 1, Right)

	if err := compare(output, expected); err != nil {
		t.Error(err)
	}
}

func TestApply(t *testing.T) {
	bms := beams{}
	bms.append(0, 1, Down)
	bms.append(1, 1, Up)
	bms.append(2, 1, Right)

	array := [][]mirrorState{
		{Empty, Empty, Empty},
		{MirrorTopLeftBottomRight, SplitHorizontal, MirrorTopLeftBottomRight},
		{Empty, Empty, Empty},
	}
	output := applyArrayToBeam(bms, array)

	expected := beams{}
	expected.append(0, 1, Right)
	expected.append(1, 1, Left)
	expected.append(1, 1, Right)
	expected.append(2, 1, Down)

	if err := compare(output, expected); err != nil {
		t.Error(err)
	}
}

// check that a beam can detect if it is stuck in a loop
// func TestLoopDetection(t *testing.T) {
// 	beams := []beamer{newBeam(1, 0, Right)}
// 	array := [][]mirrorState{
// 		{MirrorBottomLeftTopRight, Empty, MirrorTopLeftBottomRight},
// 		{Empty, Empty, Empty},
// 		{MirrorTopLeftBottomRight, Empty, MirrorBottomLeftTopRight},
// 	}
// 	for range 1000 {
// 		// if a beam is detected as stuck it is not returned
// 		beams = stepBeamPosition(beams, 3, 3)
// 		if len(beams) == 0 {
// 			break
// 		}
// 		beams = applyArrayToBeam(beams, array)
// 	}
// 	if len(beams) != 0 {
// 		t.Error("Beam could not determine that it was stuck in a loop")
// 	}
// }

func TestExample(t *testing.T) {
	if run("../example.txt", 3) != 46 {
		t.Error()
	}
}
