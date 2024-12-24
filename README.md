# Golang BMI Calculator

A simple command-line application written in Go to calculate Body Mass Index (BMI).

## Features
- Input height and weight to calculate BMI.
- Provides BMI categories (e.g., underweight, normal weight, overweight, obesity).

## Requirements
- Go (version specified in `go.mod`)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/altynbek07/golang-bmi-calculator.git
   ```

2. Navigate to the project directory:
   ```bash
   cd golang-bmi-calculator
   ```

3. Build the project:
   ```bash
   go build -o bmi_calculator
   ```

4. Run the application:
   ```bash
   ./bmi_calculator
   ```

## Usage

The application prompts the user to input their weight (in kilograms) and height (in meters) to calculate their BMI.

Example:
```bash
Введите свой рост в сантиметрах: 1.75
Введите свой вес в кг: 70

Ваш индекс массы тела: 22.86
У вас нормальный вес
```

## Author
Altynbek Kazezov

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
