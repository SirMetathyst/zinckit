// Generated by the atom tool.  DO NOT EDIT!
// source: atom_LocalScale2
package atomcommon

import "github.com/SirMetathyst/atom"

// LocalScale2Key ...
const LocalScale2Key uint = 156111880

// LocalScale2Data ...
type LocalScale2Data struct {
	X float32
	Y float32	
}

// LocalScale2Component ...
type LocalScale2Component struct {
	context atom.Context
	data map[atom.EntityID]LocalScale2Data
}

// NewLocalScale2Component ...
func NewLocalScale2Component() *LocalScale2Component {
	return &LocalScale2Component{
		data: make(map[atom.EntityID]LocalScale2Data),
	}
}

// SetContext ...
func (c *LocalScale2Component) SetContext(ctx atom.Context) {
	if c.context == nil {
		c.context = ctx
	}
}

func init() {
	x := NewLocalScale2Component()
	context := atom.Default().RegisterComponent(LocalScale2Key, x)
	x.SetContext(context)
}

// EntityDeleted ...
func (c *LocalScale2Component) EntityDeleted(id atom.EntityID) {
	delete(c.data, id)
}

// HasEntity ...
func (c *LocalScale2Component) HasEntity(id atom.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// SetLocalScale2 ...
func (c *LocalScale2Component) SetLocalScale2(id atom.EntityID, localscale2 LocalScale2Data) {
	if c.context.HasEntity(id) {
		if c.HasEntity(id) {
			c.data[id] = localscale2
			c.context.ComponentUpdated(LocalScale2Key, id)
		} else {
			c.data[id] = localscale2
			c.context.ComponentAdded(LocalScale2Key, id)
		}
	}
}

// LocalScale2 ...
func (c *LocalScale2Component) LocalScale2(id atom.EntityID) LocalScale2Data {
	return c.data[id]
}

// DeleteLocalScale2 ...
func (c *LocalScale2Component) DeleteLocalScale2(id atom.EntityID) {
	delete(c.data, id)
	c.context.ComponentDeleted(LocalScale2Key, id)
}

// SetLocalScale2X ...
func SetLocalScale2X(e *atom.EntityManager, id atom.EntityID, localscale2 LocalScale2Data) {
	v, _ := e.Component(LocalScale2Key)
	c := v.(*LocalScale2Component)
	c.SetLocalScale2(id, localscale2)
}

// SetLocalScale2 ...
func SetLocalScale2(id atom.EntityID, localscale2 LocalScale2Data) {
	SetLocalScale2X(atom.Default(), id, localscale2)
}

// LocalScale2X ...
func LocalScale2X(e *atom.EntityManager, id atom.EntityID) LocalScale2Data {
	v, _ := e.Component(LocalScale2Key)
	c := v.(*LocalScale2Component)
	return c.LocalScale2(id)
}

// LocalScale2 ...
func LocalScale2(id atom.EntityID) LocalScale2Data {
	return LocalScale2X(atom.Default(), id)
}

// DeleteLocalScale2X ...
func DeleteLocalScale2X(e *atom.EntityManager, id atom.EntityID) {
	v, _ := e.Component(LocalScale2Key)
	c := v.(*LocalScale2Component)
	c.DeleteLocalScale2(id)
}

// DeleteLocalScale2 ...
func DeleteLocalScale2(id atom.EntityID) {
	DeleteLocalScale2X(atom.Default(), id)
}

// HasLocalScale2X ...
func HasLocalScale2X(e *atom.EntityManager, id atom.EntityID) bool {
	v, _ := e.Component(LocalScale2Key)
	return v.HasEntity(id)
}

// HasLocalScale2 ...
func HasLocalScale2(id atom.EntityID) bool {
	return HasLocalScale2X(atom.Default(), id)
}