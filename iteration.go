package iteration

const repeatedCount = 5

func Repeat(character string) (repeated string) {
  for i := 0; i < repeatedCount; i++ {
    repeated += character
  }

  return
}
