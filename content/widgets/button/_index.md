+++
title = "Button"
date = 2024-10-04T19:45:29+03:00
menuPre = '<i class="icon-button"></i> '
weight = 1
draft = true
+++

A simple button with text that can be set to a background color or image for each state and a callback to react to events.

<!--more-->

<!-- ![preview](examples/wid_but_pre.png)

{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_pre.txt" lang="go" id="root" >}}
{{% /expand %}}

### Widget options

###### Text padding

Responsible for setting multiline text on a button.

{{< tabs title="Padding:" style="transparent" color="#131a22ff" >}}

{{% tab title="Left" %}}
{{< code src="assets/examples/wid_but_tex_pad_lef.txt" lang="go" id="this" >}}
![lef](examples/wid_but_tex_pad_lef.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pad_lef.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Right" %}}
{{< code src="assets/examples/wid_but_tex_pad_rig.txt" lang="go" id="this" >}}
![rig](examples/wid_but_tex_pad_rig.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pad_rig.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Top" %}}
{{< code src="assets/examples/wid_but_tex_pad_top.txt" lang="go" id="this" >}}
![top](examples/wid_but_tex_pad_top.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pad_top.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Bottom" %}}
{{< code src="assets/examples/wid_but_tex_pad_bot.txt" lang="go" id="this" >}}
![bot](examples/wid_but_tex_pad_bot.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pad_bot.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Text position

Responsible for setting multiline text on a button.

{{< tabs title="Position:" style="transparent" color="#131a22ff" >}}

{{% tab title="StartxStart" %}}
{{< code src="assets/examples/lay_gri_pos_staxsta.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_pos_staxsta.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_staxsta.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="CenterxCenter" %}}
{{< code src="assets/examples/lay_gri_pos_cenxcen.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_pos_cenxcen.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_cenxcen.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="EndxEnd" %}}
{{< code src="assets/examples/lay_gri_pos_endxend.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_pos_endxend.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_endxend.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="CenterxStart" %}}
{{< code src="assets/examples/lay_gri_pos_cenxsta.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_pos_cenxsta.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_cenxsta.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="StartxCenter" %}}
{{< code src="assets/examples/lay_gri_pos_staxcen.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_pos_staxcen.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_staxcen.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="EndxStart" %}}
{{< code src="assets/examples/lay_gri_pos_endxsta.txt" lang="go" id="this" >}}
![hei](examples/lay_gri_pos_endxsta.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_endxsta.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="StartxEnd" %}}
{{< code src="assets/examples/lay_gri_pos_staxend.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_pos_staxend.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_staxend.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Text Label

Responsible for setting multiline text on a button.

{{< tabs title="Label:" style="transparent" color="#131a22ff" >}}

{{% tab title="Login" %}}
{{< code src="assets/examples/wid_but_lab_log.txt" lang="go" id="this" >}}
![log](examples/wid_but_lab_log.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_lab_log.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Signup" %}}
{{< code src="assets/examples/wid_but_lab_sig.txt" lang="go" id="this" >}}
![horz](examples/wid_but_lab_sig.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_lab_sig.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Text Face

Responsible for setting the font family and font size.

{{< tabs title="Font:" style="transparent" color="#131a22ff" >}}

{{% tab title="Goregular-64" %}}
{{< code src="assets/examples/wid_but_fon_g64.txt" lang="go" id="this" >}}
![g64](examples/wid_but_fon_g64.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_fon_g64.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Basicfont-32" %}}
{{< code src="assets/examples/wid_but_fon_b32.txt" lang="go" id="this" >}}
![horz](examples/wid_but_fon_b32.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_fon_b32.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Text Color

Responsible for setting the font family and font size.

{{< tabs title="Font:" style="transparent" color="#131a22ff" >}}

{{% tab title="Goregular-64" %}}
{{< code src="assets/examples/wid_but_fon_g64.txt" lang="go" id="this" >}}
![g64](examples/wid_but_fon_g64.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_fon_g64.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Basicfont-32" %}}
{{< code src="assets/examples/wid_but_fon_b32.txt" lang="go" id="this" >}}
![horz](examples/wid_but_fon_b32.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_fon_b32.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Image

Responsible for the background of the button; it can be filled with color or image tiles.

{{< tabs title="Font:" style="transparent" color="#131a22ff" >}}

{{% tab title="Color" %}}
{{< code src="assets/examples/wid_but_img_col.txt" lang="go" id="this" >}}
![col](examples/wid_but_img_col.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_img_col.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Tiles" %}}
{{< code src="assets/examples/wid_but_img_til.txt" lang="go" id="this" >}}
![horz](examples/wid_but_img_til.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_img_til.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}} -->