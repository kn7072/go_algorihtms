
//https://stackoverflow.com/questions/28800672/how-to-add-new-methods-to-an-existing-type-in-go/28800807#28800807

/*
Method 1 - Type Definition

type child parent
// or
type MyThing imported.Thing

Provides access to the fields.
Does not provide access to the methods.


Method 2 - Embedding (official documentation)

type child struct {
    parent
}
// or with import and pointer
type MyThing struct {
    *imported.Thing
}


Provides access to the fields.
Provides access to the methods.
Requires consideration for initialization.


Summary

-   Using the composition method the embedded parent will not initialize if it is a pointer. The parent must be initialized separately.
-   If the embedded parent is a pointer and is not initialized when the child is initialized, a nil pointer dereference error will occur.
-   Both type definition and embed cases provide access to the fields of the parent.
-   The type definition does not allow access to the parent's methods, but embedding the parent does.

*/

package main

import (
    "fmt"
)

type parent struct {
    attr string
}

type childAlias parent

type childObjParent struct {
    parent
}

type childPointerParent struct {
    *parent
}

func (p *parent) parentDo(s string) { fmt.Println(s) }
func (c *childAlias) childAliasDo(s string) { fmt.Println(s) }
func (c *childObjParent) childObjParentDo(s string) { fmt.Println(s) }
func (c *childPointerParent) childPointerParentDo(s string) { fmt.Println(s) }

func main() {
    p := &parent{"pAttr"}
    c1 := &childAlias{"cAliasAttr"}
    c2 := &childObjParent{}
    // When the parent is a pointer it must be initialized.
    // Otherwise, we get a nil pointer error when trying to set the attr.
    c3 := &childPointerParent{}
    c4 := &childPointerParent{&parent{}}

    c2.attr = "cObjParentAttr"
    // c3.attr = "cPointerParentAttr" // NOGO nil pointer dereference
    c4.attr = "cPointerParentAttr"

    // CAN do because we inherit parent's fields
    fmt.Println(p.attr)
    fmt.Println(c1.attr)
    fmt.Println(c2.attr)
    fmt.Println(c4.attr)

    p.parentDo("called parentDo on parent")
    c1.childAliasDo("called childAliasDo on ChildAlias")
    c2.childObjParentDo("called childObjParentDo on ChildObjParent")
    c3.childPointerParentDo("called childPointerParentDo on ChildPointerParent")
    c4.childPointerParentDo("called childPointerParentDo on ChildPointerParent")

    // CANNOT do because we don't inherit parent's methods
    // c1.parentDo("called parentDo on childAlias") // NOGO c1.parentDo undefined

    // CAN do because we inherit the parent's methods
    c2.parentDo("called parentDo on childObjParent")
    c3.parentDo("called parentDo on childPointerParent")
    c4.parentDo("called parentDo on childPointerParent")
}