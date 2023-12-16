package tools

// pa push the top first element of stack b to stack a
func (a *Stack) Pa(b *Stack) {
	if len(b.items) > 0 {
		topB := b.items[0]
		a.Shift(topB)
		b.items = b.items[1:]
	}
}

// pb push the top first element of stack a to stack b
func (b *Stack) Pb(a *Stack) {
	if len(a.items) > 0 {
		topA := a.items[0]
		b.Shift(topA)
		a.items = a.items[1:]
	}
}

// sa swap first 2 elements of stack a
func (a *Stack) Sa() {
	if len(a.items) > 1 {
		a.items[0], a.items[1] = a.items[1], a.items[0]
	}
}

// sb swap first 2 elements of stack b
func (b *Stack) Sb() {
	if len(b.items) > 1 {
		b.items[0], b.items[1] = b.items[1], b.items[0]
	}
}

// ss execute sa and sb
func (a *Stack) Ss(b *Stack) {
	a.Sa()
	b.Sb()
}

// ra rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
func (a *Stack) Ra() {
	if len(a.items) > 1 {
		first := a.items[0]
		a.items = append(a.items[1:], first)
	}
}

// rb rotate stack b
func (b *Stack) Rb() {
	if len(b.items) > 1 {
		first := b.items[0]
		b.items = append(b.items[1:], first)
	}
}

// rr execute ra and rb
func (a *Stack) Rr(b *Stack) {
	a.Ra()
	b.Rb()
}

// rra reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
func (a *Stack) Rra() {
	if len(a.items) > 1 {
		last := a.items[len(a.items)-1]
		res := a.items[:len(a.items)-1]
		a.items = append([]int{last}, res...)
	}
}

// rrb reverse rotate b
func (b *Stack) Rrb() {
	if len(b.items) > 1 {
		last := b.items[len(b.items)-1]
		b.items = append([]int{last}, b.items[:len(b.items)-1]...)
	}
}

// rrr execute rra and rrb
func (a *Stack) Rrr(b *Stack) {
	a.Rra()
	b.Rrb()
}
