+++
title = "Anchor"
date = 2024-10-04T19:45:29+03:00
menuPre = '<i class="icon-anchor"></i> '
weight = 1
+++

Anchor layout allows the child containers to be anchored to a specific constraint.

<!--more-->

![preview](examples/lay_anc_pre.png)

{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pre.txt" lang="go" id="root" >}}
{{% /expand %}}

### Layout options

###### Padding

Layout allows you to specify padding for all child elements but not the itself.

{{< tabs title="Padding:" style="transparent" color="#131a22ff" >}}

{{% tab title="Left" %}}
{{< code src="assets/examples/lay_anc_pad_lef.txt" lang="go" id="this" >}}
![left](examples/lay_anc_pad_lef.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pad_lef.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Right" %}}
{{< code src="assets/examples/lay_anc_pad_rig.txt" lang="go" id="this" >}}
![left](examples/lay_anc_pad_rig.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pad_rig.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Top" %}}
{{< code src="assets/examples/lay_anc_pad_top.txt" lang="go" id="this" >}}
![left](examples/lay_anc_pad_top.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pad_top.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Bottom" %}}
{{< code src="assets/examples/lay_anc_pad_bot.txt" lang="go" id="this" >}}
![left](examples/lay_anc_pad_bot.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pad_bot.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

### Layout data

###### Horizontal position

Responsible for aligning the element along the horizontal axis.

{{< tabs title="Position:" style="transparent" color="#131a22ff" >}}

{{% tab title="Start" %}}
{{< code src="assets/examples/lay_anc_hor_pos_sta.txt" lang="go" id="this" >}}
![center](examples/lay_anc_hor_pos_sta.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_hor_pos_sta.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Center" %}}
{{< code src="assets/examples/lay_anc_hor_pos_cen.txt" lang="go" id="this" >}}
![center](examples/lay_anc_hor_pos_cen.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_hor_pos_cen.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="End" %}}
{{< code src="assets/examples/lay_anc_hor_pos_end.txt" lang="go" id="this" >}}
![center](examples/lay_anc_hor_pos_end.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_hor_pos_end.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Vertical position

Responsible for aligning the element along the vertical axis.

{{< tabs title="Position:" style="transparent" color="#131a22ff" >}}

{{% tab title="Start" %}}
{{< code src="assets/examples/lay_anc_ver_pos_sta.txt" lang="go" id="this" >}}
![center](examples/lay_anc_ver_pos_sta.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_ver_pos_sta.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Center" %}}
{{< code src="assets/examples/lay_anc_ver_pos_cen.txt" lang="go" id="this" >}}
![center](examples/lay_anc_ver_pos_cen.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_ver_pos_cen.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="End" %}}
{{< code src="assets/examples/lay_anc_ver_pos_end.txt" lang="go" id="this" >}}
![center](examples/lay_anc_ver_pos_end.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_ver_pos_end.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}
