package dao

import "fyne.io/fyne/v2"

type Diagonal struct {
}

func (d *Diagonal) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()
		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}
func (d *Diagonal) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, containerSize.Height - d.MinSize(objects).Height)
	for _, o := range objects {
		size := o.MinSize()
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(size.Width, size.Height))
	}
}