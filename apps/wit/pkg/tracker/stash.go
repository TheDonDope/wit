package tracker

// Stasher provides a way to add a given amount to the current stash
type Stasher interface {
	// Add increases the current stash inventory by the given amount.
	Add(amount float64)
}

// Grinder provides a way to remove a given amount from the current stash
type Grinder interface {
	// Consume retrieves the given amount from the current stash.
	Consume(amount float64)
}

// Stash holds the information about the inventory of a specific strain.
type Stash struct {
	Strain string  `json:"strain"`
	Amount float64 `json:"amount"`
}
