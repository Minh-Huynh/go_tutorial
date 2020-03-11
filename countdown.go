package main
import (
  "io"
  "fmt"
  "os"
  "time"
)

const FinalWord = "GO!"
const countdownStart = 3

type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep() {
  time.Sleep(1 * time.Second)
}

func main(){
  sleeper := &DefaultSleeper{}
  Countdown(os.Stdout,sleeper)
}

type Sleeper interface {
  Sleep()
}

type SpySleeper struct {
  Calls int
}
//Remember, we have a pointer passed in because Go does pass-by-copy
func (s *SpySleeper) Sleep() {
  s.Calls++
}

func Countdown(out io.Writer,sleep Sleeper){
  for i := countdownStart; i > 0; i-- {
    sleep.Sleep()
    fmt.Fprintln(out, i)
  }
  sleep.Sleep()
  fmt.Fprint(out, FinalWord)
}
