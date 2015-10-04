package alioss
import (
    "testing"
    "time"
)

func tTestISO8601Time(t *testing.T) {
    m := time.Now()
    n := ISO8601Time{m.UTC()}
    s1, _ := m.MarshalText()
    s2, _ := n.MarshalText()
    t.Log(string(s1), "\n")
    t.Log(string(s2), "\n")
}