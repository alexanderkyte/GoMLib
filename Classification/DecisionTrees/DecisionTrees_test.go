package GoMLib

import(
  "testing"
)

func TestShan(t *testing.T) {
  tempShan := shannonEntropy([]string{"yes", "yes", "no", "no", "no"})
  if tempShan > 1 || tempShan < 0.9 {
    t.Errorf("Shannon algorithm failed test")
  }
  tempShan := shannonEntropy([]string{"maybe", "yes", "no", "no", "no"})
  if tempShan > 1.4 || tempShan < 1.2 {
    t.Errorf("Shannon algorithm failed test")
  }
}
