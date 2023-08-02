package team

import employee "github.com/nithishgwd/employee_Module" // Update with your module path

// Team represents a team in the office.
type Team struct {
	name           string
	employees      []*employee.Employee
	employeesToday int // Number of employees present on a particular day
}

// NewTeam creates a new instance of Team with the given name.
func NewTeam(name string) *Team {
	return &Team{
		name:           name,
		employees:      []*employee.Employee{},
		employeesToday: 0,
	}
}

// GetName returns the name of the team.
func (t *Team) GetName() string {
	return t.name
}

// GetEmployees returns the slice of employees in the team.
func (t *Team) GetEmployees() []*employee.Employee {
	return t.employees
}

// GetEmployeesToday returns the number of employees present today in the team.
func (t *Team) GetEmployeesToday() int {
	return t.employeesToday
}

// AddEmployee adds an employee to the team.
func (t *Team) AddEmployee(emp *employee.Employee) {
	t.employees = append(t.employees, emp)
}

// UpdateEmployeesToday updates the number of employees present today in the team.
func (t *Team) UpdateEmployeesToday(numEmployees int) {
	t.employeesToday = numEmployees
}
