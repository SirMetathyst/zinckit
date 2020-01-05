// Package zinckit ...
// Generated by the zinc tool.  DO NOT EDIT!
// Source: zinc_LocalPosition2
package zinckit

import (
	"github.com/SirMetathyst/zinc"

)

// LocalPosition2Key ...
const LocalPosition2Key uint = 1388677675

// LocalPosition2Data ...
type LocalPosition2Data struct {
	X float32
	Y float32	
}

// LocalPosition2Component ...
type LocalPosition2Component struct {
	ctx zinc.CTX
	data map[zinc.EntityID]LocalPosition2Data
}

// NewLocalPosition2Component ...
func NewLocalPosition2Component() *LocalPosition2Component {
	return &LocalPosition2Component{
		data: make(map[zinc.EntityID]LocalPosition2Data),
	}
}

// SetContext ...
func (c *LocalPosition2Component) SetContext(ctx zinc.CTX) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

func init() {
	x := NewLocalPosition2Component()
	ctx := zinc.Default().RegisterComponent(LocalPosition2Key, x)
	x.SetContext(ctx)
}

// DeleteEntity ...
func (c *LocalPosition2Component) DeleteEntity(id zinc.EntityID) {
	delete(c.data, id)
}

// HasEntity ...
func (c *LocalPosition2Component) HasEntity(id zinc.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// SetLocalPosition2 ...
func (c *LocalPosition2Component) SetLocalPosition2(id zinc.EntityID, localposition2 LocalPosition2Data) {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			c.data[id] = localposition2
			c.ctx.ComponentUpdated(LocalPosition2Key, id)
		} else {
			c.data[id] = localposition2
			c.ctx.ComponentAdded(LocalPosition2Key, id)
		}
	}
}

// LocalPosition2 ...
func (c *LocalPosition2Component) LocalPosition2(id zinc.EntityID) LocalPosition2Data {
	return c.data[id]
}

// DeleteLocalPosition2 ...
func (c *LocalPosition2Component) DeleteLocalPosition2(id zinc.EntityID) {
	delete(c.data, id)
	c.ctx.ComponentDeleted(LocalPosition2Key, id)
}

// SetLocalPosition2X ...
func SetLocalPosition2X(e *zinc.EntityManager, id zinc.EntityID, localposition2 LocalPosition2Data) {
	v, _ := e.Component(LocalPosition2Key)
	c := v.(*LocalPosition2Component)
	c.SetLocalPosition2(id, localposition2)
}

// SetLocalPosition2 ...
func SetLocalPosition2(id zinc.EntityID, localposition2 LocalPosition2Data) {
	SetLocalPosition2X(zinc.Default(), id, localposition2)
}

// LocalPosition2X ...
func LocalPosition2X(e *zinc.EntityManager, id zinc.EntityID) LocalPosition2Data {
	v, _ := e.Component(LocalPosition2Key)
	c := v.(*LocalPosition2Component)
	return c.LocalPosition2(id)
}

// LocalPosition2 ...
func LocalPosition2(id zinc.EntityID) LocalPosition2Data {
	return LocalPosition2X(zinc.Default(), id)
}

// DeleteLocalPosition2X ...
func DeleteLocalPosition2X(e *zinc.EntityManager, id zinc.EntityID) {
	v, _ := e.Component(LocalPosition2Key)
	c := v.(*LocalPosition2Component)
	c.DeleteLocalPosition2(id)
}

// DeleteLocalPosition2 ...
func DeleteLocalPosition2(id zinc.EntityID) {
	DeleteLocalPosition2X(zinc.Default(), id)
}

// HasLocalPosition2X ...
func HasLocalPosition2X(e *zinc.EntityManager, id zinc.EntityID) bool {
	v, _ := e.Component(LocalPosition2Key)
	return v.HasEntity(id)
}

// HasLocalPosition2 ...
func HasLocalPosition2(id zinc.EntityID) bool {
	return HasLocalPosition2X(zinc.Default(), id)
}