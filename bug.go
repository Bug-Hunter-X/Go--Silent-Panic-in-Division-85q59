func MyFunc(a, b int) (int, error) {
  if a == 0 {
    return 0, errors.New("a cannot be zero")
  }
  return a / b, nil
}

func main() {
  result, err := MyFunc(10, 0)
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("Result:", result)
  }
}