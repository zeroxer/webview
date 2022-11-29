package webview

/*
#cgo linux openbsd freebsd netbsd CXXFLAGS: -DWEBVIEW_GTK -std=c++11
#cgo linux openbsd freebsd netbsd pkg-config: gtk+-3.0 webkit2gtk-4.0

#cgo darwin CXXFLAGS: -DWEBVIEW_COCOA -std=c++11
#cgo darwin LDFLAGS: -framework WebKit

#cgo windows CXXFLAGS: -DWEBVIEW_EDGE -std=c++17
#cgo windows LDFLAGS: -static -lole32 -lshell32 -lshlwapi -luser32

#include "webview.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

func init() {
	// Ensure that main.main is called from the main thread
	runtime.LockOSThread()
}

// Hints are used to configure window sizing and resizing
type Hint int

const (
	// Width and height are default size
	HintNone = C.WEBVIEW_HINT_NONE

	// Window size can not be changed by a user
	HintFixed = C.WEBVIEW_HINT_FIXED

	// Width and height are minimum bounds
	HintMin = C.WEBVIEW_HINT_MIN

	// Width and height are maximum bounds
	HintMax = C.WEBVIEW_HINT_MAX
)

type WebView interface {

	// Run runs the main loop until it's terminated. After this function exits -
	// you must destroy the webview.
	Run()

	// Terminate stops the main loop. It is safe to call this function from
	// a background thread.
	Terminate()

	// Dispatch posts a function to be executed on the main thread. You normally
	// do not need to call this function, unless you want to tweak the native
	// window.
	Dispatch(f func())

	// Destroy destroys a webview and closes the native window.
	Destroy()

	// Window returns a native window handle pointer. When using GTK backend the
	// pointer is GtkWindow pointer, when using Cocoa backend the pointer is
	// NSWindow pointer, when using Win32 backend the pointer is HWND pointer.
	Window() unsafe.Pointer

	// SetTitle updates the title of the native window. Must be called from the UI
	// thread.
	SetTitle(title string)

	// SetSize updates native window size. See Hint constants.
	SetSize(w int, h int, hint Hint)

	// Navigate navigates webview to the given URL. URL may be a properly encoded data.
	// URI. Examples:
	// w.Navigate("https://github.com/webview/webview")
	// w.Navigate("data:text/html,%3Ch1%3EHello%3C%2Fh1%3E")
	// w.Navigate("data:text/html;base64,PGgxPkhlbGxvPC9oMT4=")
	Navigate(url string)

	// SetHtml sets the webview HTML directly.
	// Example: w.SetHtml(w, "<h1>Hello</h1>");
	SetHtml(html string)

	// Init injects JavaScript code at the initialization of the new page. Every
	// time the webview will open a the new page - this initialization code will
	// be executed. It is guaranteed that code is executed before window.onload.
	Init(js string)

	// Eval evaluates arbitrary JavaScript code. Evaluation happens asynchronously,
	// also the result of the expression is ignored. Use RPC bindings if you want
	// to receive notifications about the results of the evaluation.
	Eval(js string)

	// Bind binds a callback function so that it will appear under the given name
	// as a global JavaScript function. Internally it uses webview_init().
	// Callback receives a request string and a user-provided argument pointer.
	// Request string is a JSON array of all the arguments passed to the
	// JavaScript function.
	//
	// f must be a function
	// f must return either value and error or just error
	Bind(name string, f interface{}) error
}

type webview struct {
	w C.webview_t
}

// New calls NewWindow to create a new window and a new webview instance. If debug
// is non-zero - developer tools will be enabled (if the platform supports them).
func New(debug bool) WebView { return NewWindow(debug, nil) }

// NewWindow creates a new webview instance. If debug is non-zero - developer
// tools will be enabled (if the platform supports them). Window parameter can be
// a pointer to the native window handle. If it's non-null - then child WebView is
// embedded into the given parent window. Otherwise a new window is created.
// Depending on the platform, a GtkWindow, NSWindow or HWND pointer can be passed
// here. Returns nil on failure. Creation can fail for various reasons such as when
// required runtime dependencies are missing or when window creation fails.
func NewWindow(debug bool, window unsafe.Pointer) WebView {
	return &webview{w: nil}
}

func (w *webview) Destroy() {
}

func (w *webview) Run() {
}

func (w *webview) Terminate() {
}

func (w *webview) Window() unsafe.Pointer {
	return nil
}

func (w *webview) Navigate(url string) {
}

func (w *webview) SetHtml(html string) {
}

func (w *webview) SetTitle(title string) {
}

func (w *webview) SetSize(width int, height int, hint Hint) {
}

func (w *webview) Init(js string) {
}

func (w *webview) Eval(js string) {
}

func (w *webview) Dispatch(f func()) {
}

func (w *webview) Bind(name string, f interface{}) error {
	return nil
}
