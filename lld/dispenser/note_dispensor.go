package dispenser

import "fmt"

type NoteDispenser struct {
	denomination int
	count        int
	next         Dispenser
}

func NewNoteDispenser(denomination, count int) *NoteDispenser {
	return &NoteDispenser{
		denomination: denomination,
		count:        count,
	}
}

func (d *NoteDispenser) SetNext(next Dispenser) {
	d.next = next
}

func (d *NoteDispenser) Dispense(amount int) error {
	if amount <= 0 {
		return nil
	}

	if amount >= d.denomination && d.count > 0 {
		numOfNotes := min(amount/d.denomination, d.count)
		dispensedAmount := numOfNotes * d.denomination
		remaining := amount - dispensedAmount

		if numOfNotes > 0 {
			fmt.Printf("Dispensing: %d x Rs.%d notes\n", numOfNotes, d.denomination)
			d.count -= numOfNotes
		}

		if remaining > 0 {
			if d.next != nil {
				return d.next.Dispense(remaining)
			}
			return fmt.Errorf("cannot dispense rs.%d - insufficient denomination", amount)
		}
	}

	if d.next != nil {
		return d.next.Dispense(amount)
	}

	return fmt.Errorf("cannot dispense rs.%d - insufficient denomination", amount)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
