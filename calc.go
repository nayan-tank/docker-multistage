package main


import (
		"fmt"
	)

	func add(x, y float64) float64 {
			return x + y
		}

		func subtract(x, y float64) float64 {
				return x - y
			}

			func multiply(x, y float64) float64 {
					return x * y
				}

				func divide(x, y float64) float64 {
						return x / y
					}

					func main() {
							var operator string
								var num1, num2 float64

									fmt.Println("Enter an operator (+, -, *, /): ")
										fmt.Scanln(&operator)

											fmt.Println("Enter two numbers: ")
												fmt.Scanln(&num1)
													fmt.Scanln(&num2)

														switch operator {
																case "+":
																			fmt.Printf("%f + %f = %f", num1, num2, add(num1, num2))
																				case "-":
																							fmt.Printf("%f - %f = %f", num1, num2, subtract(num1, num2))
																								case "*":
																											fmt.Printf("%f * %f = %f", num1, num2, multiply(num1, num2))
																												case "/":
																															fmt.Printf("%f / %f = %f", num1, num2, divide(num1, num2))
																																default:
																																			fmt.Println("Invalid operator")
																																				}
																																			}
