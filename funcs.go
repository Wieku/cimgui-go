package cimgui

// #include "cimgui_wrapper.h"
import "C"
import "unsafe"

func DrawListSplitter_SetCurrentChannel(self *ImDrawListSplitter, draw_list *ImDrawList, channel_idx int32) {
	C.DrawListSplitter_SetCurrentChannel((*C.ImDrawListSplitter)(self), (*C.ImDrawList)(draw_list), C.int(channel_idx))
}

func DrawList_PrimWriteIdx(self *ImDrawList, idx ImDrawIdx) {
	C.DrawList_PrimWriteIdx((*C.ImDrawList)(self), C.ImDrawIdx(idx))
}

func DebugCheckVersionAndDataLayout(version_str string, sz_io uint64, sz_style uint64, sz_vec2 uint64, sz_vec4 uint64, sz_drawvert uint64, sz_drawidx uint64) bool {
	version_strArg, version_strFin := wrapString(version_str)
	defer version_strFin()

	return C.DebugCheckVersionAndDataLayout(version_strArg, C.ulong(sz_io), C.ulong(sz_style), C.ulong(sz_vec2), C.ulong(sz_vec4), C.ulong(sz_drawvert), C.ulong(sz_drawidx)) != C.bool(true)
}

func OpenPopupOnItemClick(str_id string, popup_flags ImGuiPopupFlags) {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	C.OpenPopupOnItemClick(str_idArg, C.ImGuiPopupFlags(popup_flags))
}

func PopTextWrapPos() {
	C.PopTextWrapPos()
}

func GetMainViewport() *ImGuiViewport {
	return (*ImGuiViewport)(C.GetMainViewport())
}

func Font_SetGlyphVisible(self *ImFont, c ImWchar, visible bool) {
	visibleArg := C.bool(visible)

	C.Font_SetGlyphVisible((*C.ImFont)(self), C.ImWchar(c), visibleArg)
}

func DragInt(label string, v *int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapInt32(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt(labelArg, vArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func IO_AddFocusEvent(self *ImGuiIO, focused bool) {
	focusedArg := C.bool(focused)

	C.IO_AddFocusEvent((*C.ImGuiIO)(self), focusedArg)
}

func SliderFloat4(label string, v *[4]float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat4(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func DragFloat(label string, v *float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapFloat(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat(labelArg, vArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetMousePos(pOut *ImVec2) {
	C.GetMousePos((*C.ImVec2)(pOut))
}

func GetTextLineHeightWithSpacing() float32 {
	return float32(C.GetTextLineHeightWithSpacing())
}

func DrawList_AddRectFilled(self *ImDrawList, p_min ImVec2, p_max ImVec2, col uint32, rounding float32, flags ImDrawFlags) {
	C.DrawList_AddRectFilled((*C.ImDrawList)(self), C.ImVec2(p_min), C.ImVec2(p_max), C.uint(col), C.float(rounding), C.ImDrawFlags(flags))
}

func FontAtlas_GetGlyphRangesChineseFull(self *ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesChineseFull((*C.ImFontAtlas)(self)))
}

func BeginListBox(label string, size ImVec2) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.BeginListBox(labelArg, C.ImVec2(size)) != C.bool(true)
}

func FindViewportByID(id ImGuiID) *ImGuiViewport {
	return (*ImGuiViewport)(C.FindViewportByID(C.ImGuiID(id)))
}

func DrawListSplitter_Clear(self *ImDrawListSplitter) {
	C.DrawListSplitter_Clear((*C.ImDrawListSplitter)(self))
}

func FindViewportByPlatformHandle(platform_handle unsafe.Pointer) *ImGuiViewport {
	return (*ImGuiViewport)(C.FindViewportByPlatformHandle(platform_handle))
}

func Font_FindGlyphNoFallback(self *ImFont, c ImWchar) *ImFontGlyph {
	return (*ImFontGlyph)(C.Font_FindGlyphNoFallback((*C.ImFont)(self), C.ImWchar(c)))
}

func MemFree(ptr unsafe.Pointer) {
	C.MemFree(ptr)
}

func SetCursorScreenPos(pos ImVec2) {
	C.SetCursorScreenPos(C.ImVec2(pos))
}

func IO_ClearInputCharacters(self *ImGuiIO) {
	C.IO_ClearInputCharacters((*C.ImGuiIO)(self))
}

func SetNextWindowFocus() {
	C.SetNextWindowFocus()
}

func DrawList_ChannelsSetCurrent(self *ImDrawList, n int32) {
	C.DrawList_ChannelsSetCurrent((*C.ImDrawList)(self), C.int(n))
}

func Bullet() {
	C.Bullet()
}

func CalcTextSize(pOut *ImVec2, text string, text_end string, hide_text_after_double_hash bool, wrap_width float32) {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	hide_text_after_double_hashArg := C.bool(hide_text_after_double_hash)

	C.CalcTextSize((*C.ImVec2)(pOut), textArg, text_endArg, hide_text_after_double_hashArg, C.float(wrap_width))
}

func PushButtonRepeat(repeat bool) {
	repeatArg := C.bool(repeat)

	C.PushButtonRepeat(repeatArg)
}

func Spacing() {
	C.Spacing()
}

func GetMousePosOnOpeningCurrentPopup(pOut *ImVec2) {
	C.GetMousePosOnOpeningCurrentPopup((*C.ImVec2)(pOut))
}

func SetNextWindowSize(size ImVec2, cond ImGuiCond) {
	C.SetNextWindowSize(C.ImVec2(size), C.ImGuiCond(cond))
}

func Font_Font() *ImFont {
	return (*ImFont)(C.Font_Font())
}

func IsItemVisible() bool {
	return C.IsItemVisible() != C.bool(true)
}

func IO_AddInputCharacter(self *ImGuiIO, c uint32) {
	C.IO_AddInputCharacter((*C.ImGuiIO)(self), C.uint(c))
}

func TableGetRowIndex() int {
	return int(C.TableGetRowIndex())
}

func FontAtlas_ClearFonts(self *ImFontAtlas) {
	C.FontAtlas_ClearFonts((*C.ImFontAtlas)(self))
}

func Font_GetCharAdvance(self *ImFont, c ImWchar) float32 {
	return float32(C.Font_GetCharAdvance((*C.ImFont)(self), C.ImWchar(c)))
}

func GetTextLineHeight() float32 {
	return float32(C.GetTextLineHeight())
}

func RenderPlatformWindowsDefault(platform_render_arg unsafe.Pointer, renderer_render_arg unsafe.Pointer) {
	C.RenderPlatformWindowsDefault(platform_render_arg, renderer_render_arg)
}

func EndTabBar() {
	C.EndTabBar()
}

func PopID() {
	C.PopID()
}

func SliderInt3(label string, v *[3]int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt3(labelArg, vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func TableSetBgColor(target ImGuiTableBgTarget, color uint32, column_n int32) {
	C.TableSetBgColor(C.ImGuiTableBgTarget(target), C.uint(color), C.int(column_n))
}

func TextFilter_Draw(self *ImGuiTextFilter, label string, width float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.TextFilter_Draw((*C.ImGuiTextFilter)(self), labelArg, C.float(width)) != C.bool(true)
}

func IsKeyDown(key ImGuiKey) bool {
	return C.IsKeyDown(C.ImGuiKey(key)) != C.bool(true)
}

func ProgressBar(fraction float32, size_arg ImVec2, overlay string) {
	overlayArg, overlayFin := wrapString(overlay)
	defer overlayFin()

	C.ProgressBar(C.float(fraction), C.ImVec2(size_arg), overlayArg)
}

func AcceptDragDropPayload(typeArg string, flags ImGuiDragDropFlags) *ImGuiPayload {
	typeArgArg, typeArgFin := wrapString(typeArg)
	defer typeArgFin()

	return (*ImGuiPayload)(C.AcceptDragDropPayload(typeArgArg, C.ImGuiDragDropFlags(flags)))
}

func Storage_SetBool(self *ImGuiStorage, key ImGuiID, val bool) {
	valArg := C.bool(val)

	C.Storage_SetBool((*C.ImGuiStorage)(self), C.ImGuiID(key), valArg)
}

func DrawList_AddPolyline(self *ImDrawList, points *ImVec2, num_points int32, col uint32, flags ImDrawFlags, thickness float32) {
	C.DrawList_AddPolyline((*C.ImDrawList)(self), (*C.ImVec2)(points), C.int(num_points), C.uint(col), C.ImDrawFlags(flags), C.float(thickness))
}

func FontAtlas_Build(self *ImFontAtlas) bool {
	return C.FontAtlas_Build((*C.ImFontAtlas)(self)) != C.bool(true)
}

func InputTextCallbackData_ClearSelection(self *ImGuiInputTextCallbackData) {
	C.InputTextCallbackData_ClearSelection((*C.ImGuiInputTextCallbackData)(self))
}

func TableSortSpecs_TableSortSpecs() *ImGuiTableSortSpecs {
	return (*ImGuiTableSortSpecs)(C.TableSortSpecs_TableSortSpecs())
}

func DrawList_AddRect(self *ImDrawList, p_min ImVec2, p_max ImVec2, col uint32, rounding float32, flags ImDrawFlags, thickness float32) {
	C.DrawList_AddRect((*C.ImDrawList)(self), C.ImVec2(p_min), C.ImVec2(p_max), C.uint(col), C.float(rounding), C.ImDrawFlags(flags), C.float(thickness))
}

func GetMouseCursor() ImGuiMouseCursor {
	return ImGuiMouseCursor(C.GetMouseCursor())
}

func GetWindowWidth() float32 {
	return float32(C.GetWindowWidth())
}

func ColorConvertRGBtoHSV(r float32, g float32, b float32, out_h *float32, out_s *float32, out_v *float32) {
	out_hArg, out_hFin := wrapFloat(out_h)
	defer out_hFin()

	out_sArg, out_sFin := wrapFloat(out_s)
	defer out_sFin()

	out_vArg, out_vFin := wrapFloat(out_v)
	defer out_vFin()

	C.ColorConvertRGBtoHSV(C.float(r), C.float(g), C.float(b), out_hArg, out_sArg, out_vArg)
}

func GetContentRegionMax(pOut *ImVec2) {
	C.GetContentRegionMax((*C.ImVec2)(pOut))
}

func LogFinish() {
	C.LogFinish()
}

func FontAtlas_GetGlyphRangesKorean(self *ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesKorean((*C.ImFontAtlas)(self)))
}

func Payload_Payload() *ImGuiPayload {
	return (*ImGuiPayload)(C.Payload_Payload())
}

func DrawList_GetClipRectMin(pOut *ImVec2, self *ImDrawList) {
	C.DrawList_GetClipRectMin((*C.ImVec2)(pOut), (*C.ImDrawList)(self))
}

func FontGlyphRangesBuilder_FontGlyphRangesBuilder() *ImFontGlyphRangesBuilder {
	return (*ImFontGlyphRangesBuilder)(C.FontGlyphRangesBuilder_FontGlyphRangesBuilder())
}

func DrawList_PopClipRect(self *ImDrawList) {
	C.DrawList_PopClipRect((*C.ImDrawList)(self))
}

func FontAtlas_GetGlyphRangesVietnamese(self *ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesVietnamese((*C.ImFontAtlas)(self)))
}

func IsMouseReleased(button ImGuiMouseButton) bool {
	return C.IsMouseReleased(C.ImGuiMouseButton(button)) != C.bool(true)
}

func SliderAngle(label string, v_rad *float32, v_degrees_min float32, v_degrees_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	v_radArg, v_radFin := wrapFloat(v_rad)
	defer v_radFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderAngle(labelArg, v_radArg, C.float(v_degrees_min), C.float(v_degrees_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func NewFrame() {
	C.NewFrame()
}

func SetCursorPosY(local_y float32) {
	C.SetCursorPosY(C.float(local_y))
}

func VSliderFloat(label string, size ImVec2, v *float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapFloat(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.VSliderFloat(labelArg, C.ImVec2(size), vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func ColorPicker3(label string, col *[3]float32, flags ImGuiColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := (*C.float)(&col[0])

	return C.ColorPicker3(labelArg, colArg, C.ImGuiColorEditFlags(flags)) != C.bool(true)
}

func DragInt4(label string, v *[4]int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt4(labelArg, vArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func InputFloat4(label string, v *[4]float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat4(labelArg, vArg, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func TableSetupColumn(label string, flags ImGuiTableColumnFlags, init_width_or_weight float32, user_id ImGuiID) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.TableSetupColumn(labelArg, C.ImGuiTableColumnFlags(flags), C.float(init_width_or_weight), C.ImGuiID(user_id))
}

func IsMouseDoubleClicked(button ImGuiMouseButton) bool {
	return C.IsMouseDoubleClicked(C.ImGuiMouseButton(button)) != C.bool(true)
}

func ShowFontSelector(label string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.ShowFontSelector(labelArg)
}

func StyleColorsLight(dst *ImGuiStyle) {
	C.StyleColorsLight((*C.ImGuiStyle)(dst))
}

func FontAtlas_CalcCustomRectUV(self *ImFontAtlas, rect *ImFontAtlasCustomRect, out_uv_min *ImVec2, out_uv_max *ImVec2) {
	C.FontAtlas_CalcCustomRectUV((*C.ImFontAtlas)(self), (*C.ImFontAtlasCustomRect)(rect), (*C.ImVec2)(out_uv_min), (*C.ImVec2)(out_uv_max))
}

func FontAtlas_Clear(self *ImFontAtlas) {
	C.FontAtlas_Clear((*C.ImFontAtlas)(self))
}

func ShowUserGuide() {
	C.ShowUserGuide()
}

func InputTextCallbackData_InsertChars(self *ImGuiInputTextCallbackData, pos int32, text string, text_end string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	C.InputTextCallbackData_InsertChars((*C.ImGuiInputTextCallbackData)(self), C.int(pos), textArg, text_endArg)
}

func Begin(name string, p_open *bool, flags ImGuiWindowFlags) bool {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	return C.Begin(nameArg, p_openArg, C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func GetColumnsCount() int {
	return int(C.GetColumnsCount())
}

func DestroyPlatformWindows() {
	C.DestroyPlatformWindows()
}

func GetItemRectMin(pOut *ImVec2) {
	C.GetItemRectMin((*C.ImVec2)(pOut))
}

func GetPlatformIO() *ImGuiPlatformIO {
	return (*ImGuiPlatformIO)(C.GetPlatformIO())
}

func PushFont(font *ImFont) {
	C.PushFont((*C.ImFont)(font))
}

func NextColumn() {
	C.NextColumn()
}

func PopFont() {
	C.PopFont()
}

func FontGlyphRangesBuilder_Clear(self *ImFontGlyphRangesBuilder) {
	C.FontGlyphRangesBuilder_Clear((*C.ImFontGlyphRangesBuilder)(self))
}

func EndDragDropSource() {
	C.EndDragDropSource()
}

func GetWindowPos(pOut *ImVec2) {
	C.GetWindowPos((*C.ImVec2)(pOut))
}

func SliderInt2(label string, v *[2]int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt2(labelArg, vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetKeyIndex(key ImGuiKey) int {
	return int(C.GetKeyIndex(C.ImGuiKey(key)))
}

func InputTextCallbackData_HasSelection(self *ImGuiInputTextCallbackData) bool {
	return C.InputTextCallbackData_HasSelection((*C.ImGuiInputTextCallbackData)(self)) != C.bool(true)
}

func GetScrollMaxY() float32 {
	return float32(C.GetScrollMaxY())
}

func SetClipboardText(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.SetClipboardText(textArg)
}

func SetNextFrameWantCaptureKeyboard(want_capture_keyboard bool) {
	want_capture_keyboardArg := C.bool(want_capture_keyboard)

	C.SetNextFrameWantCaptureKeyboard(want_capture_keyboardArg)
}

func CalcItemWidth() float32 {
	return float32(C.CalcItemWidth())
}

func Indent(indent_w float32) {
	C.Indent(C.float(indent_w))
}

func IsMouseDown(button ImGuiMouseButton) bool {
	return C.IsMouseDown(C.ImGuiMouseButton(button)) != C.bool(true)
}

func PopStyleColor(count int32) {
	C.PopStyleColor(C.int(count))
}

func DrawList_AddCircleFilled(self *ImDrawList, center ImVec2, radius float32, col uint32, num_segments int32) {
	C.DrawList_AddCircleFilled((*C.ImDrawList)(self), C.ImVec2(center), C.float(radius), C.uint(col), C.int(num_segments))
}

func FontAtlas_AddFontDefault(self *ImFontAtlas, font_cfg *ImFontConfig) *ImFont {
	return (*ImFont)(C.FontAtlas_AddFontDefault((*C.ImFontAtlas)(self), (*C.ImFontConfig)(font_cfg)))
}

func GetCursorStartPos(pOut *ImVec2) {
	C.GetCursorStartPos((*C.ImVec2)(pOut))
}

func PlatformIO_PlatformIO() *ImGuiPlatformIO {
	return (*ImGuiPlatformIO)(C.PlatformIO_PlatformIO())
}

func CreateContext(shared_font_atlas *ImFontAtlas) *ImGuiContext {
	return (*ImGuiContext)(C.CreateContext((*C.ImFontAtlas)(shared_font_atlas)))
}

func PopClipRect() {
	C.PopClipRect()
}

func BeginPopupContextWindow(str_id string, popup_flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopupContextWindow(str_idArg, C.ImGuiPopupFlags(popup_flags)) != C.bool(true)
}

func DrawData_DrawData() *ImDrawData {
	return (*ImDrawData)(C.DrawData_DrawData())
}

func FontGlyphRangesBuilder_AddRanges(self *ImFontGlyphRangesBuilder, ranges *ImWchar) {
	C.FontGlyphRangesBuilder_AddRanges((*C.ImFontGlyphRangesBuilder)(self), (*C.ImWchar)(ranges))
}

func ColorEdit4(label string, col *[4]float32, flags ImGuiColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := (*C.float)(&col[0])

	return C.ColorEdit4(labelArg, colArg, C.ImGuiColorEditFlags(flags)) != C.bool(true)
}

func DragFloat2(label string, v *[2]float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat2(labelArg, vArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetStyleColorVec4(idx ImGuiCol) *ImVec4 {
	return (*ImVec4)(C.GetStyleColorVec4(C.ImGuiCol(idx)))
}

func IsMouseHoveringRect(r_min ImVec2, r_max ImVec2, clip bool) bool {
	clipArg := C.bool(clip)

	return C.IsMouseHoveringRect(C.ImVec2(r_min), C.ImVec2(r_max), clipArg) != C.bool(true)
}

func DrawList_PathBezierCubicCurveTo(self *ImDrawList, p2 ImVec2, p3 ImVec2, p4 ImVec2, num_segments int32) {
	C.DrawList_PathBezierCubicCurveTo((*C.ImDrawList)(self), C.ImVec2(p2), C.ImVec2(p3), C.ImVec2(p4), C.int(num_segments))
}

func BeginDisabled(disabled bool) {
	disabledArg := C.bool(disabled)

	C.BeginDisabled(disabledArg)
}

func SetNextWindowBgAlpha(alpha float32) {
	C.SetNextWindowBgAlpha(C.float(alpha))
}

func DrawList_PrimQuadUV(self *ImDrawList, a ImVec2, b ImVec2, c ImVec2, d ImVec2, uv_a ImVec2, uv_b ImVec2, uv_c ImVec2, uv_d ImVec2, col uint32) {
	C.DrawList_PrimQuadUV((*C.ImDrawList)(self), C.ImVec2(a), C.ImVec2(b), C.ImVec2(c), C.ImVec2(d), C.ImVec2(uv_a), C.ImVec2(uv_b), C.ImVec2(uv_c), C.ImVec2(uv_d), C.uint(col))
}

func DrawList_PrimRectUV(self *ImDrawList, a ImVec2, b ImVec2, uv_a ImVec2, uv_b ImVec2, col uint32) {
	C.DrawList_PrimRectUV((*C.ImDrawList)(self), C.ImVec2(a), C.ImVec2(b), C.ImVec2(uv_a), C.ImVec2(uv_b), C.uint(col))
}

func InputTextCallbackData_SelectAll(self *ImGuiInputTextCallbackData) {
	C.InputTextCallbackData_SelectAll((*C.ImGuiInputTextCallbackData)(self))
}

func ColorPicker4(label string, col *[4]float32, flags ImGuiColorEditFlags, ref_col *float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := (*C.float)(&col[0])

	ref_colArg, ref_colFin := wrapFloat(ref_col)
	defer ref_colFin()

	return C.ColorPicker4(labelArg, colArg, C.ImGuiColorEditFlags(flags), ref_colArg) != C.bool(true)
}

func EndMenu() {
	C.EndMenu()
}

func IsItemDeactivated() bool {
	return C.IsItemDeactivated() != C.bool(true)
}

func TableGetColumnIndex() int {
	return int(C.TableGetColumnIndex())
}

func BeginDragDropSource(flags ImGuiDragDropFlags) bool {
	return C.BeginDragDropSource(C.ImGuiDragDropFlags(flags)) != C.bool(true)
}

func EndTable() {
	C.EndTable()
}

func GetStyleColorName(idx ImGuiCol) string {
	return C.GoString(C.GetStyleColorName(C.ImGuiCol(idx)))
}

func ListClipper_ListClipper() *ImGuiListClipper {
	return (*ImGuiListClipper)(C.ListClipper_ListClipper())
}

func InputInt4(label string, v *[4]int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	return C.InputInt4(labelArg, vArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func SetColumnOffset(column_index int32, offset_x float32) {
	C.SetColumnOffset(C.int(column_index), C.float(offset_x))
}

func FontAtlas_AddFontFromMemoryCompressedTTF(self *ImFontAtlas, compressed_font_data unsafe.Pointer, compressed_font_size int32, size_pixels float32, font_cfg *ImFontConfig, glyph_ranges *ImWchar) *ImFont {
	return (*ImFont)(C.FontAtlas_AddFontFromMemoryCompressedTTF((*C.ImFontAtlas)(self), compressed_font_data, C.int(compressed_font_size), C.float(size_pixels), (*C.ImFontConfig)(font_cfg), (*C.ImWchar)(glyph_ranges)))
}

func SetColumnWidth(column_index int32, width float32) {
	C.SetColumnWidth(C.int(column_index), C.float(width))
}

func DrawList_AddBezierQuadratic(self *ImDrawList, p1 ImVec2, p2 ImVec2, p3 ImVec2, col uint32, thickness float32, num_segments int32) {
	C.DrawList_AddBezierQuadratic((*C.ImDrawList)(self), C.ImVec2(p1), C.ImVec2(p2), C.ImVec2(p3), C.uint(col), C.float(thickness), C.int(num_segments))
}

func SetNextItemOpen(is_open bool, cond ImGuiCond) {
	is_openArg := C.bool(is_open)

	C.SetNextItemOpen(is_openArg, C.ImGuiCond(cond))
}

func EndListBox() {
	C.EndListBox()
}

func FontAtlas_GetCustomRectByIndex(self *ImFontAtlas, index int32) *ImFontAtlasCustomRect {
	return (*ImFontAtlasCustomRect)(C.FontAtlas_GetCustomRectByIndex((*C.ImFontAtlas)(self), C.int(index)))
}

func IO_AddMouseButtonEvent(self *ImGuiIO, button int32, down bool) {
	downArg := C.bool(down)

	C.IO_AddMouseButtonEvent((*C.ImGuiIO)(self), C.int(button), downArg)
}

func Font_RenderChar(self *ImFont, draw_list *ImDrawList, size float32, pos ImVec2, col uint32, c ImWchar) {
	C.Font_RenderChar((*C.ImFont)(self), (*C.ImDrawList)(draw_list), C.float(size), C.ImVec2(pos), C.uint(col), C.ImWchar(c))
}

func TableHeader(label string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.TableHeader(labelArg)
}

func DrawList_GetClipRectMax(pOut *ImVec2, self *ImDrawList) {
	C.DrawList_GetClipRectMax((*C.ImVec2)(pOut), (*C.ImDrawList)(self))
}

func FontAtlas_AddFontFromFileTTF(self *ImFontAtlas, filename string, size_pixels float32, font_cfg *ImFontConfig, glyph_ranges *ImWchar) *ImFont {
	filenameArg, filenameFin := wrapString(filename)
	defer filenameFin()

	return (*ImFont)(C.FontAtlas_AddFontFromFileTTF((*C.ImFontAtlas)(self), filenameArg, C.float(size_pixels), (*C.ImFontConfig)(font_cfg), (*C.ImWchar)(glyph_ranges)))
}

func Font_IsGlyphRangeUnused(self *ImFont, c_begin uint32, c_last uint32) bool {
	return C.Font_IsGlyphRangeUnused((*C.ImFont)(self), C.uint(c_begin), C.uint(c_last)) != C.bool(true)
}

func VSliderScalar(label string, size ImVec2, data_type ImGuiDataType, p_data unsafe.Pointer, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.VSliderScalar(labelArg, C.ImVec2(size), C.ImGuiDataType(data_type), p_data, p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func SetNextFrameWantCaptureMouse(want_capture_mouse bool) {
	want_capture_mouseArg := C.bool(want_capture_mouse)

	C.SetNextFrameWantCaptureMouse(want_capture_mouseArg)
}

func SaveIniSettingsToMemory(out_ini_size *uint64) string {
	return C.GoString(C.SaveIniSettingsToMemory((*C.ulong)(out_ini_size)))
}

func DrawList_PathLineToMergeDuplicate(self *ImDrawList, pos ImVec2) {
	C.DrawList_PathLineToMergeDuplicate((*C.ImDrawList)(self), C.ImVec2(pos))
}

func DrawList_PrimReserve(self *ImDrawList, idx_count int32, vtx_count int32) {
	C.DrawList_PrimReserve((*C.ImDrawList)(self), C.int(idx_count), C.int(vtx_count))
}

func GetStyle() *ImGuiStyle {
	return (*ImGuiStyle)(C.GetStyle())
}

func SetCurrentContext(ctx *ImGuiContext) {
	C.SetCurrentContext((*C.ImGuiContext)(ctx))
}

func InputInt3(label string, v *[3]int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	return C.InputInt3(labelArg, vArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func IO_AddMouseWheelEvent(self *ImGuiIO, wh_x float32, wh_y float32) {
	C.IO_AddMouseWheelEvent((*C.ImGuiIO)(self), C.float(wh_x), C.float(wh_y))
}

func IsItemActive() bool {
	return C.IsItemActive() != C.bool(true)
}

func TableGetColumnCount() int {
	return int(C.TableGetColumnCount())
}

func GetContentRegionAvail(pOut *ImVec2) {
	C.GetContentRegionAvail((*C.ImVec2)(pOut))
}

func IsItemFocused() bool {
	return C.IsItemFocused() != C.bool(true)
}

func ListClipper_ForceDisplayRangeByIndices(self *ImGuiListClipper, item_min int32, item_max int32) {
	C.ListClipper_ForceDisplayRangeByIndices((*C.ImGuiListClipper)(self), C.int(item_min), C.int(item_max))
}

func GetStateStorage() *ImGuiStorage {
	return (*ImGuiStorage)(C.GetStateStorage())
}

func DrawList_AddConvexPolyFilled(self *ImDrawList, points *ImVec2, num_points int32, col uint32) {
	C.DrawList_AddConvexPolyFilled((*C.ImDrawList)(self), (*C.ImVec2)(points), C.int(num_points), C.uint(col))
}

func DrawList_AddDrawCmd(self *ImDrawList) {
	C.DrawList_AddDrawCmd((*C.ImDrawList)(self))
}

func DrawList_AddNgon(self *ImDrawList, center ImVec2, radius float32, col uint32, num_segments int32, thickness float32) {
	C.DrawList_AddNgon((*C.ImDrawList)(self), C.ImVec2(center), C.float(radius), C.uint(col), C.int(num_segments), C.float(thickness))
}

func Font_AddRemapChar(self *ImFont, dst ImWchar, src ImWchar, overwrite_dst bool) {
	overwrite_dstArg := C.bool(overwrite_dst)

	C.Font_AddRemapChar((*C.ImFont)(self), C.ImWchar(dst), C.ImWchar(src), overwrite_dstArg)
}

func Storage_Clear(self *ImGuiStorage) {
	C.Storage_Clear((*C.ImGuiStorage)(self))
}

func GetColumnWidth(column_index int32) float32 {
	return float32(C.GetColumnWidth(C.int(column_index)))
}

func IsAnyItemActive() bool {
	return C.IsAnyItemActive() != C.bool(true)
}

func BeginTabBar(str_id string, flags ImGuiTabBarFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginTabBar(str_idArg, C.ImGuiTabBarFlags(flags)) != C.bool(true)
}

func IsKeyReleased(key ImGuiKey) bool {
	return C.IsKeyReleased(C.ImGuiKey(key)) != C.bool(true)
}

func SliderScalarN(label string, data_type ImGuiDataType, p_data unsafe.Pointer, components int32, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderScalarN(labelArg, C.ImGuiDataType(data_type), p_data, C.int(components), p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func IO_ClearInputKeys(self *ImGuiIO) {
	C.IO_ClearInputKeys((*C.ImGuiIO)(self))
}

func FontAtlas_GetGlyphRangesJapanese(self *ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesJapanese((*C.ImFontAtlas)(self)))
}

func PopItemWidth() {
	C.PopItemWidth()
}

func DrawList_AddQuad(self *ImDrawList, p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, col uint32, thickness float32) {
	C.DrawList_AddQuad((*C.ImDrawList)(self), C.ImVec2(p1), C.ImVec2(p2), C.ImVec2(p3), C.ImVec2(p4), C.uint(col), C.float(thickness))
}

func FontGlyphRangesBuilder_AddText(self *ImFontGlyphRangesBuilder, text string, text_end string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	C.FontGlyphRangesBuilder_AddText((*C.ImFontGlyphRangesBuilder)(self), textArg, text_endArg)
}

func BeginMenuBar() bool {
	return C.BeginMenuBar() != C.bool(true)
}

func ColorButton(desc_id string, col ImVec4, flags ImGuiColorEditFlags, size ImVec2) bool {
	desc_idArg, desc_idFin := wrapString(desc_id)
	defer desc_idFin()

	return C.ColorButton(desc_idArg, C.ImVec4(col), C.ImGuiColorEditFlags(flags), C.ImVec2(size)) != C.bool(true)
}

func TableGetColumnFlags(column_n int32) ImGuiTableColumnFlags {
	return ImGuiTableColumnFlags(C.TableGetColumnFlags(C.int(column_n)))
}

func Font_ClearOutputData(self *ImFont) {
	C.Font_ClearOutputData((*C.ImFont)(self))
}

func GetFrameHeight() float32 {
	return float32(C.GetFrameHeight())
}

func GetScrollY() float32 {
	return float32(C.GetScrollY())
}

func BeginMenu(label string, enabled bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	enabledArg := C.bool(enabled)

	return C.BeginMenu(labelArg, enabledArg) != C.bool(true)
}

func CloseCurrentPopup() {
	C.CloseCurrentPopup()
}

func GetVersion() string {
	return C.GoString(C.GetVersion())
}

func DrawListSplitter_ClearFreeMemory(self *ImDrawListSplitter) {
	C.DrawListSplitter_ClearFreeMemory((*C.ImDrawListSplitter)(self))
}

func ArrowButton(str_id string, dir ImGuiDir) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.ArrowButton(str_idArg, C.ImGuiDir(dir)) != C.bool(true)
}

func InputFloat2(label string, v *[2]float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat2(labelArg, vArg, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func IsItemClicked(mouse_button ImGuiMouseButton) bool {
	return C.IsItemClicked(C.ImGuiMouseButton(mouse_button)) != C.bool(true)
}

func DrawList_PrimRect(self *ImDrawList, a ImVec2, b ImVec2, col uint32) {
	C.DrawList_PrimRect((*C.ImDrawList)(self), C.ImVec2(a), C.ImVec2(b), C.uint(col))
}

func Storage_GetBool(self *ImGuiStorage, key ImGuiID, default_val bool) bool {
	default_valArg := C.bool(default_val)

	return C.Storage_GetBool((*C.ImGuiStorage)(self), C.ImGuiID(key), default_valArg) != C.bool(true)
}

func ColorEdit3(label string, col *[3]float32, flags ImGuiColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	colArg := (*C.float)(&col[0])

	return C.ColorEdit3(labelArg, colArg, C.ImGuiColorEditFlags(flags)) != C.bool(true)
}

func EndDisabled() {
	C.EndDisabled()
}

func BeginPopup(str_id string, flags ImGuiWindowFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopup(str_idArg, C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func IsAnyMouseDown() bool {
	return C.IsAnyMouseDown() != C.bool(true)
}

func SetDragDropPayload(typeArg string, data unsafe.Pointer, sz uint64, cond ImGuiCond) bool {
	typeArgArg, typeArgFin := wrapString(typeArg)
	defer typeArgFin()

	return C.SetDragDropPayload(typeArgArg, data, C.ulong(sz), C.ImGuiCond(cond)) != C.bool(true)
}

func TableSetColumnEnabled(column_n int32, v bool) {
	vArg := C.bool(v)

	C.TableSetColumnEnabled(C.int(column_n), vArg)
}

func DrawListSplitter_DrawListSplitter() *ImDrawListSplitter {
	return (*ImDrawListSplitter)(C.DrawListSplitter_DrawListSplitter())
}

func DrawList_PrimWriteVtx(self *ImDrawList, pos ImVec2, uv ImVec2, col uint32) {
	C.DrawList_PrimWriteVtx((*C.ImDrawList)(self), C.ImVec2(pos), C.ImVec2(uv), C.uint(col))
}

func FontAtlas_GetGlyphRangesChineseSimplifiedCommon(self *ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesChineseSimplifiedCommon((*C.ImFontAtlas)(self)))
}

func InputScalarN(label string, data_type ImGuiDataType, p_data unsafe.Pointer, components int32, p_step unsafe.Pointer, p_step_fast unsafe.Pointer, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputScalarN(labelArg, C.ImGuiDataType(data_type), p_data, C.int(components), p_step, p_step_fast, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func BeginPopupContextVoid(str_id string, popup_flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopupContextVoid(str_idArg, C.ImGuiPopupFlags(popup_flags)) != C.bool(true)
}

func SliderFloat3(label string, v *[3]float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat3(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func DrawList_AddQuadFilled(self *ImDrawList, p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, col uint32) {
	C.DrawList_AddQuadFilled((*C.ImDrawList)(self), C.ImVec2(p1), C.ImVec2(p2), C.ImVec2(p3), C.ImVec2(p4), C.uint(col))
}

func TextFilter_IsActive(self *ImGuiTextFilter) bool {
	return C.TextFilter_IsActive((*C.ImGuiTextFilter)(self)) != C.bool(true)
}

func InputFloat(label string, v *float32, step float32, step_fast float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapFloat(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat(labelArg, vArg, C.float(step), C.float(step_fast), formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func NewLine() {
	C.NewLine()
}

func SliderScalar(label string, data_type ImGuiDataType, p_data unsafe.Pointer, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderScalar(labelArg, C.ImGuiDataType(data_type), p_data, p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func FontAtlas_ClearTexData(self *ImFontAtlas) {
	C.FontAtlas_ClearTexData((*C.ImFontAtlas)(self))
}

func IsWindowCollapsed() bool {
	return C.IsWindowCollapsed() != C.bool(true)
}

func LogToTTY(auto_open_depth int32) {
	C.LogToTTY(C.int(auto_open_depth))
}

func TableSetColumnIndex(column_n int32) bool {
	return C.TableSetColumnIndex(C.int(column_n)) != C.bool(true)
}

func IsItemDeactivatedAfterEdit() bool {
	return C.IsItemDeactivatedAfterEdit() != C.bool(true)
}

func Render() {
	C.Render()
}

func DragInt3(label string, v *[3]int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt3(labelArg, vArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func GetMouseClickedCount(button ImGuiMouseButton) int {
	return int(C.GetMouseClickedCount(C.ImGuiMouseButton(button)))
}

func TabItemButton(label string, flags ImGuiTabItemFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.TabItemButton(labelArg, C.ImGuiTabItemFlags(flags)) != C.bool(true)
}

func DrawList_AddRectFilledMultiColor(self *ImDrawList, p_min ImVec2, p_max ImVec2, col_upr_left uint32, col_upr_right uint32, col_bot_right uint32, col_bot_left uint32) {
	C.DrawList_AddRectFilledMultiColor((*C.ImDrawList)(self), C.ImVec2(p_min), C.ImVec2(p_max), C.uint(col_upr_left), C.uint(col_upr_right), C.uint(col_bot_right), C.uint(col_bot_left))
}

func DragScalar(label string, data_type ImGuiDataType, p_data unsafe.Pointer, v_speed float32, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragScalar(labelArg, C.ImGuiDataType(data_type), p_data, C.float(v_speed), p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func IsItemEdited() bool {
	return C.IsItemEdited() != C.bool(true)
}

func SaveIniSettingsToDisk(ini_filename string) {
	ini_filenameArg, ini_filenameFin := wrapString(ini_filename)
	defer ini_filenameFin()

	C.SaveIniSettingsToDisk(ini_filenameArg)
}

func SetStateStorage(storage *ImGuiStorage) {
	C.SetStateStorage((*C.ImGuiStorage)(storage))
}

func DrawList_ChannelsSplit(self *ImDrawList, count int32) {
	C.DrawList_ChannelsSplit((*C.ImDrawList)(self), C.int(count))
}

func SetCursorPosX(local_x float32) {
	C.SetCursorPosX(C.float(local_x))
}

func IO_IO() *ImGuiIO {
	return (*ImGuiIO)(C.IO_IO())
}

func Payload_Clear(self *ImGuiPayload) {
	C.Payload_Clear((*C.ImGuiPayload)(self))
}

func DestroyContext(ctx *ImGuiContext) {
	C.DestroyContext((*C.ImGuiContext)(ctx))
}

func DrawList_PushClipRectFullScreen(self *ImDrawList) {
	C.DrawList_PushClipRectFullScreen((*C.ImDrawList)(self))
}

func FontAtlas_AddFontFromMemoryTTF(self *ImFontAtlas, font_data unsafe.Pointer, font_size int32, size_pixels float32, font_cfg *ImFontConfig, glyph_ranges *ImWchar) *ImFont {
	return (*ImFont)(C.FontAtlas_AddFontFromMemoryTTF((*C.ImFontAtlas)(self), font_data, C.int(font_size), C.float(size_pixels), (*C.ImFontConfig)(font_cfg), (*C.ImWchar)(glyph_ranges)))
}

func GetWindowContentRegionMax(pOut *ImVec2) {
	C.GetWindowContentRegionMax((*C.ImVec2)(pOut))
}

func Storage_SetFloat(self *ImGuiStorage, key ImGuiID, val float32) {
	C.Storage_SetFloat((*C.ImGuiStorage)(self), C.ImGuiID(key), C.float(val))
}

func DragFloat3(label string, v *[3]float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat3(labelArg, vArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func SetItemAllowOverlap() {
	C.SetItemAllowOverlap()
}

func ShowStyleSelector(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.ShowStyleSelector(labelArg) != C.bool(true)
}

func FontAtlas_AddFont(self *ImFontAtlas, font_cfg *ImFontConfig) *ImFont {
	return (*ImFont)(C.FontAtlas_AddFont((*C.ImFontAtlas)(self), (*C.ImFontConfig)(font_cfg)))
}

func GetCursorPosX() float32 {
	return float32(C.GetCursorPosX())
}

func GetDrawListSharedData() *ImDrawListSharedData {
	return (*ImDrawListSharedData)(C.GetDrawListSharedData())
}

func TextFilter_Build(self *ImGuiTextFilter) {
	C.TextFilter_Build((*C.ImGuiTextFilter)(self))
}

func Dummy(size ImVec2) {
	C.Dummy(C.ImVec2(size))
}

func TextBuffer_c_str(self *ImGuiTextBuffer) string {
	return C.GoString(C.TextBuffer_c_str((*C.ImGuiTextBuffer)(self)))
}

func DrawList_AddageRounded(self *ImDrawList, user_texture_id ImTextureID, p_min ImVec2, p_max ImVec2, uv_min ImVec2, uv_max ImVec2, col uint32, rounding float32, flags ImDrawFlags) {
	C.DrawList_AddageRounded((*C.ImDrawList)(self), C.ImTextureID(user_texture_id), C.ImVec2(p_min), C.ImVec2(p_max), C.ImVec2(uv_min), C.ImVec2(uv_max), C.uint(col), C.float(rounding), C.ImDrawFlags(flags))
}

func Style_ScaleAllSizes(self *ImGuiStyle, scale_factor float32) {
	C.Style_ScaleAllSizes((*C.ImGuiStyle)(self), C.float(scale_factor))
}

func FontAtlas_IsBuilt(self *ImFontAtlas) bool {
	return C.FontAtlas_IsBuilt((*C.ImFontAtlas)(self)) != C.bool(true)
}

func FontGlyphRangesBuilder_GetBit(self *ImFontGlyphRangesBuilder, n uint64) bool {
	return C.FontGlyphRangesBuilder_GetBit((*C.ImFontGlyphRangesBuilder)(self), C.ulong(n)) != C.bool(true)
}

func EndMenuBar() {
	C.EndMenuBar()
}

func SetNextWindowContentSize(size ImVec2) {
	C.SetNextWindowContentSize(C.ImVec2(size))
}

func DrawList_PushClipRect(self *ImDrawList, clip_rect_min ImVec2, clip_rect_max ImVec2, intersect_with_current_clip_rect bool) {
	intersect_with_current_clip_rectArg := C.bool(intersect_with_current_clip_rect)

	C.DrawList_PushClipRect((*C.ImDrawList)(self), C.ImVec2(clip_rect_min), C.ImVec2(clip_rect_max), intersect_with_current_clip_rectArg)
}

func Viewport_GetCenter(pOut *ImVec2, self *ImGuiViewport) {
	C.Viewport_GetCenter((*C.ImVec2)(pOut), (*C.ImGuiViewport)(self))
}

func LoadIniSettingsFromMemory(ini_data string, ini_size uint64) {
	ini_dataArg, ini_dataFin := wrapString(ini_data)
	defer ini_dataFin()

	C.LoadIniSettingsFromMemory(ini_dataArg, C.ulong(ini_size))
}

func GetDrawData() *ImDrawData {
	return (*ImDrawData)(C.GetDrawData())
}

func ShowMetricsWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowMetricsWindow(p_openArg)
}

func TreePop() {
	C.TreePop()
}

func Viewport_GetWorkCenter(pOut *ImVec2, self *ImGuiViewport) {
	C.Viewport_GetWorkCenter((*C.ImVec2)(pOut), (*C.ImGuiViewport)(self))
}

func Payload_IsDelivery(self *ImGuiPayload) bool {
	return C.Payload_IsDelivery((*C.ImGuiPayload)(self)) != C.bool(true)
}

func Button(label string, size ImVec2) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.Button(labelArg, C.ImVec2(size)) != C.bool(true)
}

func VSliderInt(label string, size ImVec2, v *int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapInt32(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.VSliderInt(labelArg, C.ImVec2(size), vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func BeginPopupContextItem(str_id string, popup_flags ImGuiPopupFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginPopupContextItem(str_idArg, C.ImGuiPopupFlags(popup_flags)) != C.bool(true)
}

func ShowDemoWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowDemoWindow(p_openArg)
}

func SmallButton(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	return C.SmallButton(labelArg) != C.bool(true)
}

func DrawList_PushTextureID(self *ImDrawList, texture_id ImTextureID) {
	C.DrawList_PushTextureID((*C.ImDrawList)(self), C.ImTextureID(texture_id))
}

func FontAtlas_GetGlyphRangesDefault(self *ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesDefault((*C.ImFontAtlas)(self)))
}

func TableHeadersRow() {
	C.TableHeadersRow()
}

func Storage_BuildSortByKey(self *ImGuiStorage) {
	C.Storage_BuildSortByKey((*C.ImGuiStorage)(self))
}

func TextFilter_Clear(self *ImGuiTextFilter) {
	C.TextFilter_Clear((*C.ImGuiTextFilter)(self))
}

func TableColumnSortSpecs_TableColumnSortSpecs() *ImGuiTableColumnSortSpecs {
	return (*ImGuiTableColumnSortSpecs)(C.TableColumnSortSpecs_TableColumnSortSpecs())
}

func GetClipboardText() string {
	return C.GoString(C.GetClipboardText())
}

func BeginTabItem(label string, p_open *bool, flags ImGuiTabItemFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	return C.BeginTabItem(labelArg, p_openArg, C.ImGuiTabItemFlags(flags)) != C.bool(true)
}

func GetColumnIndex() int {
	return int(C.GetColumnIndex())
}

func IsWindowDocked() bool {
	return C.IsWindowDocked() != C.bool(true)
}

func DrawList_PopTextureID(self *ImDrawList) {
	C.DrawList_PopTextureID((*C.ImDrawList)(self))
}

func Font_IsLoaded(self *ImFont) bool {
	return C.Font_IsLoaded((*C.ImFont)(self)) != C.bool(true)
}

func Style_Style() *ImGuiStyle {
	return (*ImGuiStyle)(C.Style_Style())
}

func GetWindowContentRegionMin(pOut *ImVec2) {
	C.GetWindowContentRegionMin((*C.ImVec2)(pOut))
}

func SetWindowFontScale(scale float32) {
	C.SetWindowFontScale(C.float(scale))
}

func GetScrollX() float32 {
	return float32(C.GetScrollX())
}

func Color_HSV(pOut *ImColor, h float32, s float32, v float32, a float32) {
	C.Color_HSV((*C.ImColor)(pOut), C.float(h), C.float(s), C.float(v), C.float(a))
}

func BeginDragDropTarget() bool {
	return C.BeginDragDropTarget() != C.bool(true)
}

func EndFrame() {
	C.EndFrame()
}

func IsWindowAppearing() bool {
	return C.IsWindowAppearing() != C.bool(true)
}

func IO_AddInputCharactersUTF8(self *ImGuiIO, str string) {
	strArg, strFin := wrapString(str)
	defer strFin()

	C.IO_AddInputCharactersUTF8((*C.ImGuiIO)(self), strArg)
}

func TableNextColumn() bool {
	return C.TableNextColumn() != C.bool(true)
}

func IO_SetKeyEventNativeData(self *ImGuiIO, key ImGuiKey, native_keycode int32, native_scancode int32, native_legacy_index int32) {
	C.IO_SetKeyEventNativeData((*C.ImGuiIO)(self), C.ImGuiKey(key), C.int(native_keycode), C.int(native_scancode), C.int(native_legacy_index))
}

func EndTabItem() {
	C.EndTabItem()
}

func SetKeyboardFocusHere(offset int32) {
	C.SetKeyboardFocusHere(C.int(offset))
}

func FontGlyphRangesBuilder_AddChar(self *ImFontGlyphRangesBuilder, c ImWchar) {
	C.FontGlyphRangesBuilder_AddChar((*C.ImFontGlyphRangesBuilder)(self), C.ImWchar(c))
}

func OnceUponAFrame_OnceUponAFrame() *ImGuiOnceUponAFrame {
	return (*ImGuiOnceUponAFrame)(C.OnceUponAFrame_OnceUponAFrame())
}

func BeginTable(str_id string, column int32, flags ImGuiTableFlags, outer_size ImVec2, inner_width float32) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.BeginTable(str_idArg, C.int(column), C.ImGuiTableFlags(flags), C.ImVec2(outer_size), C.float(inner_width)) != C.bool(true)
}

func Font_GetDebugName(self *ImFont) string {
	return C.GoString(C.Font_GetDebugName((*C.ImFont)(self)))
}

func Font_GrowIndex(self *ImFont, new_size int32) {
	C.Font_GrowIndex((*C.ImFont)(self), C.int(new_size))
}

func GetWindowHeight() float32 {
	return float32(C.GetWindowHeight())
}

func DrawList_ChannelsMerge(self *ImDrawList) {
	C.DrawList_ChannelsMerge((*C.ImDrawList)(self))
}

func EndCombo() {
	C.EndCombo()
}

func LoadIniSettingsFromDisk(ini_filename string) {
	ini_filenameArg, ini_filenameFin := wrapString(ini_filename)
	defer ini_filenameFin()

	C.LoadIniSettingsFromDisk(ini_filenameArg)
}

func LogToFile(auto_open_depth int32, filename string) {
	filenameArg, filenameFin := wrapString(filename)
	defer filenameFin()

	C.LogToFile(C.int(auto_open_depth), filenameArg)
}

func GetCursorPos(pOut *ImVec2) {
	C.GetCursorPos((*C.ImVec2)(pOut))
}

func SliderFloat(label string, v *float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapFloat(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func Storage_SetVoidPtr(self *ImGuiStorage, key ImGuiID, val unsafe.Pointer) {
	C.Storage_SetVoidPtr((*C.ImGuiStorage)(self), C.ImGuiID(key), val)
}

func ColorConvertHSVtoRGB(h float32, s float32, v float32, out_r *float32, out_g *float32, out_b *float32) {
	out_rArg, out_rFin := wrapFloat(out_r)
	defer out_rFin()

	out_gArg, out_gFin := wrapFloat(out_g)
	defer out_gFin()

	out_bArg, out_bFin := wrapFloat(out_b)
	defer out_bFin()

	C.ColorConvertHSVtoRGB(C.float(h), C.float(s), C.float(v), out_rArg, out_gArg, out_bArg)
}

func FontAtlas_ClearInputData(self *ImFontAtlas) {
	C.FontAtlas_ClearInputData((*C.ImFontAtlas)(self))
}

func Storage_GetFloat(self *ImGuiStorage, key ImGuiID, default_val float32) float32 {
	return float32(C.Storage_GetFloat((*C.ImGuiStorage)(self), C.ImGuiID(key), C.float(default_val)))
}

func GetFontSize() float32 {
	return float32(C.GetFontSize())
}

func GetWindowViewport() *ImGuiViewport {
	return (*ImGuiViewport)(C.GetWindowViewport())
}

func InputInt(label string, v *int32, step int32, step_fast int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapInt32(v)
	defer vFin()

	return C.InputInt(labelArg, vArg, C.int(step), C.int(step_fast), C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func IsKeyPressed(key ImGuiKey, repeat bool) bool {
	repeatArg := C.bool(repeat)

	return C.IsKeyPressed(C.ImGuiKey(key), repeatArg) != C.bool(true)
}

func DrawList_CloneOutput(self *ImDrawList) *ImDrawList {
	return (*ImDrawList)(C.DrawList_CloneOutput((*C.ImDrawList)(self)))
}

func FontAtlas_AddCustomRectRegular(self *ImFontAtlas, width int32, height int32) int {
	return int(C.FontAtlas_AddCustomRectRegular((*C.ImFontAtlas)(self), C.int(width), C.int(height)))
}

func IO_SetAppAcceptingEvents(self *ImGuiIO, accepting_events bool) {
	accepting_eventsArg := C.bool(accepting_events)

	C.IO_SetAppAcceptingEvents((*C.ImGuiIO)(self), accepting_eventsArg)
}

func GetItemRectMax(pOut *ImVec2) {
	C.GetItemRectMax((*C.ImVec2)(pOut))
}

func IO_AddKeyEvent(self *ImGuiIO, key ImGuiKey, down bool) {
	downArg := C.bool(down)

	C.IO_AddKeyEvent((*C.ImGuiIO)(self), C.ImGuiKey(key), downArg)
}

func SameLine(offset_from_start_x float32, spacing float32) {
	C.SameLine(C.float(offset_from_start_x), C.float(spacing))
}

func PopAllowKeyboardFocus() {
	C.PopAllowKeyboardFocus()
}

func DrawList_AddTriangleFilled(self *ImDrawList, p1 ImVec2, p2 ImVec2, p3 ImVec2, col uint32) {
	C.DrawList_AddTriangleFilled((*C.ImDrawList)(self), C.ImVec2(p1), C.ImVec2(p2), C.ImVec2(p3), C.uint(col))
}

func FontAtlas_AddCustomRectFontGlyph(self *ImFontAtlas, font *ImFont, id ImWchar, width int32, height int32, advance_x float32, offset ImVec2) int {
	return int(C.FontAtlas_AddCustomRectFontGlyph((*C.ImFontAtlas)(self), (*C.ImFont)(font), C.ImWchar(id), C.int(width), C.int(height), C.float(advance_x), C.ImVec2(offset)))
}

func Checkbox(label string, v *bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapBool(v)
	defer vFin()

	return C.Checkbox(labelArg, vArg) != C.bool(true)
}

func SetMouseCursor(cursor_type ImGuiMouseCursor) {
	C.SetMouseCursor(C.ImGuiMouseCursor(cursor_type))
}

func SetNextWindowPos(pos ImVec2, cond ImGuiCond, pivot ImVec2) {
	C.SetNextWindowPos(C.ImVec2(pos), C.ImGuiCond(cond), C.ImVec2(pivot))
}

func ShowAboutWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowAboutWindow(p_openArg)
}

func SetTabItemClosed(tab_or_docked_window_label string) {
	tab_or_docked_window_labelArg, tab_or_docked_window_labelFin := wrapString(tab_or_docked_window_label)
	defer tab_or_docked_window_labelFin()

	C.SetTabItemClosed(tab_or_docked_window_labelArg)
}

func BeginCombo(label string, preview_value string, flags ImGuiComboFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	preview_valueArg, preview_valueFin := wrapString(preview_value)
	defer preview_valueFin()

	return C.BeginCombo(labelArg, preview_valueArg, C.ImGuiComboFlags(flags)) != C.bool(true)
}

func DragInt2(label string, v *[2]int32, v_speed float32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragInt2(labelArg, vArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func SetNextWindowCollapsed(collapsed bool, cond ImGuiCond) {
	collapsedArg := C.bool(collapsed)

	C.SetNextWindowCollapsed(collapsedArg, C.ImGuiCond(cond))
}

func GetFontTexUvWhitePixel(pOut *ImVec2) {
	C.GetFontTexUvWhitePixel((*C.ImVec2)(pOut))
}

func DrawList_AddTriangle(self *ImDrawList, p1 ImVec2, p2 ImVec2, p3 ImVec2, col uint32, thickness float32) {
	C.DrawList_AddTriangle((*C.ImDrawList)(self), C.ImVec2(p1), C.ImVec2(p2), C.ImVec2(p3), C.uint(col), C.float(thickness))
}

func Payload_IsDataType(self *ImGuiPayload, typeArg string) bool {
	typeArgArg, typeArgFin := wrapString(typeArg)
	defer typeArgFin()

	return C.Payload_IsDataType((*C.ImGuiPayload)(self), typeArgArg) != C.bool(true)
}

func IO_AddMouseViewportEvent(self *ImGuiIO, id ImGuiID) {
	C.IO_AddMouseViewportEvent((*C.ImGuiIO)(self), C.ImGuiID(id))
}

func DragIntRange2(label string, v_current_min *int32, v_current_max *int32, v_speed float32, v_min int32, v_max int32, format string, format_max string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	v_current_minArg, v_current_minFin := wrapInt32(v_current_min)
	defer v_current_minFin()

	v_current_maxArg, v_current_maxFin := wrapInt32(v_current_max)
	defer v_current_maxFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	format_maxArg, format_maxFin := wrapString(format_max)
	defer format_maxFin()

	return C.DragIntRange2(labelArg, v_current_minArg, v_current_maxArg, C.float(v_speed), C.int(v_min), C.int(v_max), formatArg, format_maxArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func SetItemDefaultFocus() {
	C.SetItemDefaultFocus()
}

func BeginMainMenuBar() bool {
	return C.BeginMainMenuBar() != C.bool(true)
}

func DragFloat4(label string, v *[4]float32, v_speed float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragFloat4(labelArg, vArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func Font_FindGlyph(self *ImFont, c ImWchar) *ImFontGlyph {
	return (*ImFontGlyph)(C.Font_FindGlyph((*C.ImFont)(self), C.ImWchar(c)))
}

func FontAtlas_SetTexID(self *ImFontAtlas, id ImTextureID) {
	C.FontAtlas_SetTexID((*C.ImFontAtlas)(self), C.ImTextureID(id))
}

func IsAnyItemFocused() bool {
	return C.IsAnyItemFocused() != C.bool(true)
}

func InputFloat3(label string, v *[3]float32, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputFloat3(labelArg, vArg, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func DrawList_PathRect(self *ImDrawList, rect_min ImVec2, rect_max ImVec2, rounding float32, flags ImDrawFlags) {
	C.DrawList_PathRect((*C.ImDrawList)(self), C.ImVec2(rect_min), C.ImVec2(rect_max), C.float(rounding), C.ImDrawFlags(flags))
}

func PopButtonRepeat() {
	C.PopButtonRepeat()
}

func IsItemHovered(flags ImGuiHoveredFlags) bool {
	return C.IsItemHovered(C.ImGuiHoveredFlags(flags)) != C.bool(true)
}

func EndPopup() {
	C.EndPopup()
}

func WindowClass_WindowClass() *ImGuiWindowClass {
	return (*ImGuiWindowClass)(C.WindowClass_WindowClass())
}

func DragScalarN(label string, data_type ImGuiDataType, p_data unsafe.Pointer, components int32, v_speed float32, p_min unsafe.Pointer, p_max unsafe.Pointer, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.DragScalarN(labelArg, C.ImGuiDataType(data_type), p_data, C.int(components), C.float(v_speed), p_min, p_max, formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func DrawList_AddCircle(self *ImDrawList, center ImVec2, radius float32, col uint32, num_segments int32, thickness float32) {
	C.DrawList_AddCircle((*C.ImDrawList)(self), C.ImVec2(center), C.float(radius), C.uint(col), C.int(num_segments), C.float(thickness))
}

func ListClipper_Begin(self *ImGuiListClipper, items_count int32, items_height float32) {
	C.ListClipper_Begin((*C.ImGuiListClipper)(self), C.int(items_count), C.float(items_height))
}

func EndMainMenuBar() {
	C.EndMainMenuBar()
}

func PopStyleVar(count int32) {
	C.PopStyleVar(C.int(count))
}

func DrawList_PathBezierQuadraticCurveTo(self *ImDrawList, p2 ImVec2, p3 ImVec2, num_segments int32) {
	C.DrawList_PathBezierQuadraticCurveTo((*C.ImDrawList)(self), C.ImVec2(p2), C.ImVec2(p3), C.int(num_segments))
}

func SliderFloat2(label string, v *[2]float32, v_min float32, v_max float32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.float)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderFloat2(labelArg, vArg, C.float(v_min), C.float(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func TableSetupScrollFreeze(cols int32, rows int32) {
	C.TableSetupScrollFreeze(C.int(cols), C.int(rows))
}

func FontAtlas_GetGlyphRangesCyrillic(self *ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesCyrillic((*C.ImFontAtlas)(self)))
}

func ShowDebugLogWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowDebugLogWindow(p_openArg)
}

func IsMouseClicked(button ImGuiMouseButton, repeat bool) bool {
	repeatArg := C.bool(repeat)

	return C.IsMouseClicked(C.ImGuiMouseButton(button), repeatArg) != C.bool(true)
}

func Font_AddGlyph(self *ImFont, src_cfg *ImFontConfig, c ImWchar, x0 float32, y0 float32, x1 float32, y1 float32, u0 float32, v0 float32, u1 float32, v1 float32, advance_x float32) {
	C.Font_AddGlyph((*C.ImFont)(self), (*C.ImFontConfig)(src_cfg), C.ImWchar(c), C.float(x0), C.float(y0), C.float(x1), C.float(y1), C.float(u0), C.float(v0), C.float(u1), C.float(v1), C.float(advance_x))
}

func DrawList_AddageQuad(self *ImDrawList, user_texture_id ImTextureID, p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, uv1 ImVec2, uv2 ImVec2, uv3 ImVec2, uv4 ImVec2, col uint32) {
	C.DrawList_AddageQuad((*C.ImDrawList)(self), C.ImTextureID(user_texture_id), C.ImVec2(p1), C.ImVec2(p2), C.ImVec2(p3), C.ImVec2(p4), C.ImVec2(uv1), C.ImVec2(uv2), C.ImVec2(uv3), C.ImVec2(uv4), C.uint(col))
}

func SetColorEditOptions(flags ImGuiColorEditFlags) {
	C.SetColorEditOptions(C.ImGuiColorEditFlags(flags))
}

func InputTextCallbackData_DeleteChars(self *ImGuiInputTextCallbackData, pos int32, bytes_count int32) {
	C.InputTextCallbackData_DeleteChars((*C.ImGuiInputTextCallbackData)(self), C.int(pos), C.int(bytes_count))
}

func BeginChildFrame(id ImGuiID, size ImVec2, flags ImGuiWindowFlags) bool {
	return C.BeginChildFrame(C.ImGuiID(id), C.ImVec2(size), C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func IsWindowFocused(flags ImGuiFocusedFlags) bool {
	return C.IsWindowFocused(C.ImGuiFocusedFlags(flags)) != C.bool(true)
}

func GetFont() *ImFont {
	return (*ImFont)(C.GetFont())
}

func InputScalar(label string, data_type ImGuiDataType, p_data unsafe.Pointer, p_step unsafe.Pointer, p_step_fast unsafe.Pointer, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputScalar(labelArg, C.ImGuiDataType(data_type), p_data, p_step, p_step_fast, formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func SetNextItemWidth(item_width float32) {
	C.SetNextItemWidth(C.float(item_width))
}

func FontAtlasCustomRect_IsPacked(self *ImFontAtlasCustomRect) bool {
	return C.FontAtlasCustomRect_IsPacked((*C.ImFontAtlasCustomRect)(self)) != C.bool(true)
}

func GetIO() *ImGuiIO {
	return (*ImGuiIO)(C.GetIO())
}

func IsMousePosValid(mouse_pos *ImVec2) bool {
	return C.IsMousePosValid((*C.ImVec2)(mouse_pos)) != C.bool(true)
}

func PushAllowKeyboardFocus(allow_keyboard_focus bool) {
	allow_keyboard_focusArg := C.bool(allow_keyboard_focus)

	C.PushAllowKeyboardFocus(allow_keyboard_focusArg)
}

func Viewport_Viewport() *ImGuiViewport {
	return (*ImGuiViewport)(C.Viewport_Viewport())
}

func TableGetSortSpecs() *ImGuiTableSortSpecs {
	return (*ImGuiTableSortSpecs)(C.TableGetSortSpecs())
}

func DrawList_PathArcToFast(self *ImDrawList, center ImVec2, radius float32, a_min_of_12 int32, a_max_of_12 int32) {
	C.DrawList_PathArcToFast((*C.ImDrawList)(self), C.ImVec2(center), C.float(radius), C.int(a_min_of_12), C.int(a_max_of_12))
}

func DrawList_Addage(self *ImDrawList, user_texture_id ImTextureID, p_min ImVec2, p_max ImVec2, uv_min ImVec2, uv_max ImVec2, col uint32) {
	C.DrawList_Addage((*C.ImDrawList)(self), C.ImTextureID(user_texture_id), C.ImVec2(p_min), C.ImVec2(p_max), C.ImVec2(uv_min), C.ImVec2(uv_max), C.uint(col))
}

func BeginTooltip() {
	C.BeginTooltip()
}

func LogButtons() {
	C.LogButtons()
}

func ShowStackToolWindow(p_open *bool) {
	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	C.ShowStackToolWindow(p_openArg)
}

func SetScrollHereY(center_y_ratio float32) {
	C.SetScrollHereY(C.float(center_y_ratio))
}

func InputTextCallbackData_InputTextCallbackData() *ImGuiInputTextCallbackData {
	return (*ImGuiInputTextCallbackData)(C.InputTextCallbackData_InputTextCallbackData())
}

func EndGroup() {
	C.EndGroup()
}

func IsWindowHovered(flags ImGuiHoveredFlags) bool {
	return C.IsWindowHovered(C.ImGuiHoveredFlags(flags)) != C.bool(true)
}

func DrawData_DeIndexAllBuffers(self *ImDrawData) {
	C.DrawData_DeIndexAllBuffers((*C.ImDrawData)(self))
}

func DrawList_PathLineTo(self *ImDrawList, pos ImVec2) {
	C.DrawList_PathLineTo((*C.ImDrawList)(self), C.ImVec2(pos))
}

func InputDouble(label string, v *float64, step float64, step_fast float64, format string, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.InputDouble(labelArg, (*C.double)(v), C.double(step), C.double(step_fast), formatArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func GetKeyName(key ImGuiKey) string {
	return C.GoString(C.GetKeyName(C.ImGuiKey(key)))
}

func Storage_GetInt(self *ImGuiStorage, key ImGuiID, default_val int32) int {
	return int(C.Storage_GetInt((*C.ImGuiStorage)(self), C.ImGuiID(key), C.int(default_val)))
}

func PushTextWrapPos(wrap_local_pos_x float32) {
	C.PushTextWrapPos(C.float(wrap_local_pos_x))
}

func SetNextWindowDockID(dock_id ImGuiID, cond ImGuiCond) {
	C.SetNextWindowDockID(C.ImGuiID(dock_id), C.ImGuiCond(cond))
}

func GetCursorScreenPos(pOut *ImVec2) {
	C.GetCursorScreenPos((*C.ImVec2)(pOut))
}

func DrawList_AddLine(self *ImDrawList, p1 ImVec2, p2 ImVec2, col uint32, thickness float32) {
	C.DrawList_AddLine((*C.ImDrawList)(self), C.ImVec2(p1), C.ImVec2(p2), C.uint(col), C.float(thickness))
}

func SliderInt4(label string, v *[4]int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt4(labelArg, vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func IO_AddKeyAnalogEvent(self *ImGuiIO, key ImGuiKey, down bool, v float32) {
	downArg := C.bool(down)

	C.IO_AddKeyAnalogEvent((*C.ImGuiIO)(self), C.ImGuiKey(key), downArg, C.float(v))
}

func BeginPopupModal(name string, p_open *bool, flags ImGuiWindowFlags) bool {
	nameArg, nameFin := wrapString(name)
	defer nameFin()

	p_openArg, p_openFin := wrapBool(p_open)
	defer p_openFin()

	return C.BeginPopupModal(nameArg, p_openArg, C.ImGuiWindowFlags(flags)) != C.bool(true)
}

func DebugTextEncoding(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.DebugTextEncoding(textArg)
}

func FontAtlas_GetGlyphRangesThai(self *ImFontAtlas) *ImWchar {
	return (*ImWchar)(C.FontAtlas_GetGlyphRangesThai((*C.ImFontAtlas)(self)))
}

func PlatformMonitor_PlatformMonitor() *ImGuiPlatformMonitor {
	return (*ImGuiPlatformMonitor)(C.PlatformMonitor_PlatformMonitor())
}

func EndTooltip() {
	C.EndTooltip()
}

func GetCurrentContext() *ImGuiContext {
	return (*ImGuiContext)(C.GetCurrentContext())
}

func GetDragDropPayload() *ImGuiPayload {
	return (*ImGuiPayload)(C.GetDragDropPayload())
}

func SetScrollHereX(center_x_ratio float32) {
	C.SetScrollHereX(C.float(center_x_ratio))
}

func TextFilter_PassFilter(self *ImGuiTextFilter, text string, text_end string) bool {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	return C.TextFilter_PassFilter((*C.ImGuiTextFilter)(self), textArg, text_endArg) != C.bool(true)
}

func Font_RenderText(self *ImFont, draw_list *ImDrawList, size float32, pos ImVec2, col uint32, clip_rect ImVec4, text_begin string, text_end string, wrap_width float32, cpu_fine_clip bool) {
	text_beginArg, text_beginFin := wrapString(text_begin)
	defer text_beginFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	cpu_fine_clipArg := C.bool(cpu_fine_clip)

	C.Font_RenderText((*C.ImFont)(self), (*C.ImDrawList)(draw_list), C.float(size), C.ImVec2(pos), C.uint(col), C.ImVec4(clip_rect), text_beginArg, text_endArg, C.float(wrap_width), cpu_fine_clipArg)
}

func EndDragDropTarget() {
	C.EndDragDropTarget()
}

func GetFrameHeightWithSpacing() float32 {
	return float32(C.GetFrameHeightWithSpacing())
}

func ListClipper_Step(self *ImGuiListClipper) bool {
	return C.ListClipper_Step((*C.ImGuiListClipper)(self)) != C.bool(true)
}

func Storage_SetAllInt(self *ImGuiStorage, val int32) {
	C.Storage_SetAllInt((*C.ImGuiStorage)(self), C.int(val))
}

func GetTreeNodeToLabelSpacing() float32 {
	return float32(C.GetTreeNodeToLabelSpacing())
}

func Storage_SetInt(self *ImGuiStorage, key ImGuiID, val int32) {
	C.Storage_SetInt((*C.ImGuiStorage)(self), C.ImGuiID(key), C.int(val))
}

func BeginGroup() {
	C.BeginGroup()
}

func ShowStyleEditor(ref *ImGuiStyle) {
	C.ShowStyleEditor((*C.ImGuiStyle)(ref))
}

func DrawListSplitter_Merge(self *ImDrawListSplitter, draw_list *ImDrawList) {
	C.DrawListSplitter_Merge((*C.ImDrawListSplitter)(self), (*C.ImDrawList)(draw_list))
}

func TextBuffer_TextBuffer() *ImGuiTextBuffer {
	return (*ImGuiTextBuffer)(C.TextBuffer_TextBuffer())
}

func FontAtlas_FontAtlas() *ImFontAtlas {
	return (*ImFontAtlas)(C.FontAtlas_FontAtlas())
}

func Font_BuildLookupTable(self *ImFont) {
	C.Font_BuildLookupTable((*C.ImFont)(self))
}

func Payload_IsPreview(self *ImGuiPayload) bool {
	return C.Payload_IsPreview((*C.ImGuiPayload)(self)) != C.bool(true)
}

func LogToClipboard(auto_open_depth int32) {
	C.LogToClipboard(C.int(auto_open_depth))
}

func DrawListSplitter_Split(self *ImDrawListSplitter, draw_list *ImDrawList, count int32) {
	C.DrawListSplitter_Split((*C.ImDrawListSplitter)(self), (*C.ImDrawList)(draw_list), C.int(count))
}

func DrawList_AddBezierCubic(self *ImDrawList, p1 ImVec2, p2 ImVec2, p3 ImVec2, p4 ImVec2, col uint32, thickness float32, num_segments int32) {
	C.DrawList_AddBezierCubic((*C.ImDrawList)(self), C.ImVec2(p1), C.ImVec2(p2), C.ImVec2(p3), C.ImVec2(p4), C.uint(col), C.float(thickness), C.int(num_segments))
}

func DrawList_PrimUnreserve(self *ImDrawList, idx_count int32, vtx_count int32) {
	C.DrawList_PrimUnreserve((*C.ImDrawList)(self), C.int(idx_count), C.int(vtx_count))
}

func GetMouseDragDelta(pOut *ImVec2, button ImGuiMouseButton, lock_threshold float32) {
	C.GetMouseDragDelta((*C.ImVec2)(pOut), C.ImGuiMouseButton(button), C.float(lock_threshold))
}

func InputInt2(label string, v *[2]int32, flags ImGuiInputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg := (*C.int)(&v[0])

	return C.InputInt2(labelArg, vArg, C.ImGuiInputTextFlags(flags)) != C.bool(true)
}

func TextUnformatted(text string, text_end string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	C.TextUnformatted(textArg, text_endArg)
}

func DrawData_Clear(self *ImDrawData) {
	C.DrawData_Clear((*C.ImDrawData)(self))
}

func DragFloatRange2(label string, v_current_min *float32, v_current_max *float32, v_speed float32, v_min float32, v_max float32, format string, format_max string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	v_current_minArg, v_current_minFin := wrapFloat(v_current_min)
	defer v_current_minFin()

	v_current_maxArg, v_current_maxFin := wrapFloat(v_current_max)
	defer v_current_maxFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	format_maxArg, format_maxFin := wrapString(format_max)
	defer format_maxFin()

	return C.DragFloatRange2(labelArg, v_current_minArg, v_current_maxArg, C.float(v_speed), C.float(v_min), C.float(v_max), formatArg, format_maxArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func InvisibleButton(str_id string, size ImVec2, flags ImGuiButtonFlags) bool {
	str_idArg, str_idFin := wrapString(str_id)
	defer str_idFin()

	return C.InvisibleButton(str_idArg, C.ImVec2(size), C.ImGuiButtonFlags(flags)) != C.bool(true)
}

func SetNextWindowViewport(viewport_id ImGuiID) {
	C.SetNextWindowViewport(C.ImGuiID(viewport_id))
}

func GetCursorPosY() float32 {
	return float32(C.GetCursorPosY())
}

func GetFrameCount() int {
	return int(C.GetFrameCount())
}

func SetNextWindowClass(window_class *ImGuiWindowClass) {
	C.SetNextWindowClass((*C.ImGuiWindowClass)(window_class))
}

func GetWindowSize(pOut *ImVec2) {
	C.GetWindowSize((*C.ImVec2)(pOut))
}

func FontGlyphRangesBuilder_SetBit(self *ImFontGlyphRangesBuilder, n uint64) {
	C.FontGlyphRangesBuilder_SetBit((*C.ImFontGlyphRangesBuilder)(self), C.ulong(n))
}

func GetItemRectSize(pOut *ImVec2) {
	C.GetItemRectSize((*C.ImVec2)(pOut))
}

func Unindent(indent_w float32) {
	C.Unindent(C.float(indent_w))
}

func DrawData_ScaleClipRects(self *ImDrawData, fb_scale ImVec2) {
	C.DrawData_ScaleClipRects((*C.ImDrawData)(self), C.ImVec2(fb_scale))
}

func IsMouseDragging(button ImGuiMouseButton, lock_threshold float32) bool {
	return C.IsMouseDragging(C.ImGuiMouseButton(button), C.float(lock_threshold)) != C.bool(true)
}

func GetColumnOffset(column_index int32) float32 {
	return float32(C.GetColumnOffset(C.int(column_index)))
}

func IsAnyItemHovered() bool {
	return C.IsAnyItemHovered() != C.bool(true)
}

func UpdatePlatformWindows() {
	C.UpdatePlatformWindows()
}

func SetCursorPos(local_pos ImVec2) {
	C.SetCursorPos(C.ImVec2(local_pos))
}

func DrawList_PathFillConvex(self *ImDrawList, col uint32) {
	C.DrawList_PathFillConvex((*C.ImDrawList)(self), C.uint(col))
}

func TableNextRow(row_flags ImGuiTableRowFlags, min_row_height float32) {
	C.TableNextRow(C.ImGuiTableRowFlags(row_flags), C.float(min_row_height))
}

func FontAtlas_AddFontFromMemoryCompressedBase85TTF(self *ImFontAtlas, compressed_font_data_base85 string, size_pixels float32, font_cfg *ImFontConfig, glyph_ranges *ImWchar) *ImFont {
	compressed_font_data_base85Arg, compressed_font_data_base85Fin := wrapString(compressed_font_data_base85)
	defer compressed_font_data_base85Fin()

	return (*ImFont)(C.FontAtlas_AddFontFromMemoryCompressedBase85TTF((*C.ImFontAtlas)(self), compressed_font_data_base85Arg, C.float(size_pixels), (*C.ImFontConfig)(font_cfg), (*C.ImWchar)(glyph_ranges)))
}

func GetWindowDpiScale() float32 {
	return float32(C.GetWindowDpiScale())
}

func DrawList_PrimVtx(self *ImDrawList, pos ImVec2, uv ImVec2, col uint32) {
	C.DrawList_PrimVtx((*C.ImDrawList)(self), C.ImVec2(pos), C.ImVec2(uv), C.uint(col))
}

func GetKeyPressedAmount(key ImGuiKey, repeat_delay float32, rate float32) int {
	return int(C.GetKeyPressedAmount(C.ImGuiKey(key), C.float(repeat_delay), C.float(rate)))
}

func SliderInt(label string, v *int32, v_min int32, v_max int32, format string, flags ImGuiSliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	vArg, vFin := wrapInt32(v)
	defer vFin()

	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.SliderInt(labelArg, vArg, C.int(v_min), C.int(v_max), formatArg, C.ImGuiSliderFlags(flags)) != C.bool(true)
}

func PushItemWidth(item_width float32) {
	C.PushItemWidth(C.float(item_width))
}

func Columns(count int32, id string, border bool) {
	idArg, idFin := wrapString(id)
	defer idFin()

	borderArg := C.bool(border)

	C.Columns(C.int(count), idArg, borderArg)
}

func End() {
	C.End()
}

func DrawList_AddNgonFilled(self *ImDrawList, center ImVec2, radius float32, col uint32, num_segments int32) {
	C.DrawList_AddNgonFilled((*C.ImDrawList)(self), C.ImVec2(center), C.float(radius), C.uint(col), C.int(num_segments))
}

func ColorConvertU32ToFloat4(pOut *ImVec4, in uint32) {
	C.ColorConvertU32ToFloat4((*C.ImVec4)(pOut), C.uint(in))
}

func EndChild() {
	C.EndChild()
}

func DrawCmd_DrawCmd() *ImDrawCmd {
	return (*ImDrawCmd)(C.DrawCmd_DrawCmd())
}

func DrawList_PathArcTo(self *ImDrawList, center ImVec2, radius float32, a_min float32, a_max float32, num_segments int32) {
	C.DrawList_PathArcTo((*C.ImDrawList)(self), C.ImVec2(center), C.float(radius), C.float(a_min), C.float(a_max), C.int(num_segments))
}

func Font_CalcWordWrapPositionA(self *ImFont, scale float32, text string, text_end string, wrap_width float32) string {
	textArg, textFin := wrapString(text)
	defer textFin()

	text_endArg, text_endFin := wrapString(text_end)
	defer text_endFin()

	return C.GoString(C.Font_CalcWordWrapPositionA((*C.ImFont)(self), C.float(scale), textArg, text_endArg, C.float(wrap_width)))
}

func ListClipper_End(self *ImGuiListClipper) {
	C.ListClipper_End((*C.ImGuiListClipper)(self))
}

func GetScrollMaxX() float32 {
	return float32(C.GetScrollMaxX())
}

func StyleColorsDark(dst *ImGuiStyle) {
	C.StyleColorsDark((*C.ImGuiStyle)(dst))
}

func IsItemToggledOpen() bool {
	return C.IsItemToggledOpen() != C.bool(true)
}

func IO_AddMousePosEvent(self *ImGuiIO, x float32, y float32) {
	C.IO_AddMousePosEvent((*C.ImGuiIO)(self), C.float(x), C.float(y))
}

func EndChildFrame() {
	C.EndChildFrame()
}

func PushClipRect(clip_rect_min ImVec2, clip_rect_max ImVec2, intersect_with_current_clip_rect bool) {
	intersect_with_current_clip_rectArg := C.bool(intersect_with_current_clip_rect)

	C.PushClipRect(C.ImVec2(clip_rect_min), C.ImVec2(clip_rect_max), intersect_with_current_clip_rectArg)
}

func FontAtlasCustomRect_FontAtlasCustomRect() *ImFontAtlasCustomRect {
	return (*ImFontAtlasCustomRect)(C.FontAtlasCustomRect_FontAtlasCustomRect())
}

func AlignTextToFramePadding() {
	C.AlignTextToFramePadding()
}

func StyleColorsClassic(dst *ImGuiStyle) {
	C.StyleColorsClassic((*C.ImGuiStyle)(dst))
}

func Color_SetHSV(self *ImColor, h float32, s float32, v float32, a float32) {
	C.Color_SetHSV((*C.ImColor)(self), C.float(h), C.float(s), C.float(v), C.float(a))
}

func TextFilter_TextFilter(default_filter string) *ImGuiTextFilter {
	default_filterArg, default_filterFin := wrapString(default_filter)
	defer default_filterFin()

	return (*ImGuiTextFilter)(C.TextFilter_TextFilter(default_filterArg))
}

func GetWindowDrawList() *ImDrawList {
	return (*ImDrawList)(C.GetWindowDrawList())
}

func DrawList_PathStroke(self *ImDrawList, col uint32, flags ImDrawFlags, thickness float32) {
	C.DrawList_PathStroke((*C.ImDrawList)(self), C.uint(col), C.ImDrawFlags(flags), C.float(thickness))
}

func DrawList_DrawList(shared_data *ImDrawListSharedData) *ImDrawList {
	return (*ImDrawList)(C.DrawList_DrawList((*C.ImDrawListSharedData)(shared_data)))
}

func Separator() {
	C.Separator()
}

func DrawList_PathClear(self *ImDrawList) {
	C.DrawList_PathClear((*C.ImDrawList)(self))
}

func IsItemActivated() bool {
	return C.IsItemActivated() != C.bool(true)
}

func ResetMouseDragDelta(button ImGuiMouseButton) {
	C.ResetMouseDragDelta(C.ImGuiMouseButton(button))
}
