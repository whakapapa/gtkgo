// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18,!gtk_3_20

package gtk

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
)

// PopupAtPointer() is a wrapper for gtk_menu_popup_at_pointer()
func (v *Menu) PopupAtPointer(triggerEvent *gdk.Event) {
	e := (*C.GdkEvent)(unsafe.Pointer(triggerEvent.Native()))
	C.gtk_menu_popup_at_pointer(v.native(), e)
}

// PopupAtWidget() is a wrapper for gtk_menu_popup_at_widget()
func (v *Menu) PopupAtWidget(widget IWidget, widgetAnchor gdk.GdkGravity, menuAnchor gdk.GdkGravity, triggerEvent *gdk.Event) {
	e := (*C.GdkEvent)(unsafe.Pointer(triggerEvent.Native()))
	C.gtk_menu_popup_at_widget(v.native(), widget.toWidget(), C.GdkGravity(widgetAnchor), C.GdkGravity(menuAnchor), e)
}
