package main

// User represents the top-level app user's data
type User struct {
	ID        string `bson:"_id"`
	Email     string `bson:"email"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
}

// Attribute represents an attribute key/value pair that can
// be attached to another object
type Attribute struct {
	Key   string `bson:"key"`   // User-defined attribute key
	Value string `bson:"value"` // User-defined attribute value
}

// Organization represents an organization defined by a user
// (and specific to that user).
type Organization struct {
	ID       string      `bson:"_id"`      // Organization / Object ID
	Owner    string      `bson:"owner"`    // User ID of the owner
	Name     string      `bson:"name"`     // Name of the organization
	Location string      `bson:"location"` // Location of the organization
	Attrs    []Attribute `bson:"attrs"`    // Additional, user-defined attributes
}

// Person represents a person defined by a user (and specific to that user).
type Person struct {
	ID         string `bson:"_id"`   // Person / Object ID
	Owner      string `bson:"owner"` // User ID of the owner
	FirstName  string `bson:"firstName"`
	LastName   string `bson:"lastName"`
	MiddleName string `bson:"middleName,omitempty"`
	NickName   string `bson:"nickName,omitempty"`
	Location   string `bson:"location,omitempty"`
	Phone      string `bson:"phone,omitempty"`
	Email      string `bson:"email,omitempty"`
	Address    struct {
		Street  string `bson:"street,omitempty"`
		City    string `bson:"city,omitempty"`
		State   string `bson:"state,omitempty"`
		Country string `bson:"country,omitempty"`
		ZipCode string `bson:"zipCode,omitempty"`
	} `bson:"address,omitempty"`
	JobHistory []struct {
		CompanyID string      `bson:"companyID"` // IDs referring to the company where they work
		Title     string      `bson:"title"`     // The user's title at the job
		StartDate string      `bson:"startDate"`
		EndDate   string      `bson:"endDate,omitempty"`
		Attrs     []Attribute `bson:"attrs"` // Additional, user-defined attributes
	} `bson:"jobHistory"` // The user's job history
	Attrs []Attribute `bson:"attrs"` // Additional, user-defined attributes
}

// Event represents an event defined by a user (and specific to that user).
// An event can be associated with a person, an organization, or both.
// Some examples of possible events include a meeting, a club event,
// grabbing coffee, a birthday party, etc.
type Event struct {
	ID    string      `bson:"_id"`   // Event / Object ID
	Owner string      `bson:"owner"` // User ID of the owner
	Name  string      `bson:"name"`  // Name of the event
	Date  string      `bson:"date"`  // Date of the event
	Notes string      `bson:"notes"` // Notes about the event
	Attrs []Attribute `bson:"attrs"` // Additional, user-defined attributes
}

// JobApplication represents a job application defined by a user (and specific to that user).
// A job application can be associated with a person, an organization, or both.
type JobApplication struct {
	ID          string      `bson:"_id"`         // JobApplication / Object ID
	Owner       string      `bson:"owner"`       // User ID of the owner
	JobTitle    string      `bson:"jobTitle"`    // The job application title
	CompanyID   string      `bson:"companyID"`   // ID of the company the job application is for
	CompanyName string      `bson:"companyName"` // Name of the company the job application is for
	Notes       string      `bson:"notes"`       // Notes about the job application
	EventIDs    []string    `bson:"eventIDs"`    // IDs of the events associated with the job application
	People      []string    `bson:"people"`      // IDs of the people associated with the job application (eg hiring managers, recruiters, other applicants, etc)
	Attrs       []Attribute `bson:"attrs"`       // Additional, user-defined attributes
}
