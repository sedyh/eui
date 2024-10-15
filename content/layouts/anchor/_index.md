+++
title = "Anchor layout"
date = 2024-10-04T19:45:29+03:00
weight = 1
+++

Anchor layout allows the child containers to be anchored to a specific constraint.

![preview](examples/lay_anc_pre.png)

{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pre.txt" lang="go" id="full" >}}
{{% /expand %}}

### Layout options

###### Padding

Layout allows you to specify padding for all child elements but not the itself.

{{< tabs title="Padding:" style="transparent" color="#131a22ff" >}}
{{% tab title="Left" %}}
{{< code src="assets/examples/lay_anc_pad_lef.txt" lang="go" id="pad" >}}
![left](examples/lay_anc_pad_lef.png)
{{% /tab %}}
{{% tab title="Right" %}}
{{< code src="assets/examples/lay_anc_pad_rig.txt" lang="go" id="pad" >}}
![left](examples/lay_anc_pad_rig.png)
{{% /tab %}}
{{% tab title="Top" %}}
{{< code src="assets/examples/lay_anc_pad_top.txt" lang="go" id="pad" >}}
![left](examples/lay_anc_pad_top.png)
{{% /tab %}}
{{% tab title="Bottom" %}}
{{< code src="assets/examples/lay_anc_pad_bot.txt" lang="go" id="pad" >}}
![left](examples/lay_anc_pad_bot.png)
{{% /tab %}}
{{< /tabs >}}

### Layout data

###### Horizontal position

Responsible for aligning the element along the horizontal axis.

{{< tabs title="Position:" style="transparent" color="#131a22ff" >}}
{{% tab title="Start" %}}
{{< code src="assets/examples/lay_anc_hor_pos_sta.txt" lang="go" id="pos" >}}
![center](examples/lay_anc_hor_pos_sta.png)
{{% /tab %}}
{{% tab title="Center" %}}
{{< code src="assets/examples/lay_anc_hor_pos_cen.txt" lang="go" id="pos" >}}
![center](examples/lay_anc_hor_pos_cen.png)
{{% /tab %}}
{{% tab title="End" %}}
{{< code src="assets/examples/lay_anc_hor_pos_end.txt" lang="go" id="pos" >}}
![center](examples/lay_anc_hor_pos_end.png)
{{% /tab %}}
{{< /tabs >}}

###### Vertical position

Responsible for aligning the element along the vertical axis.

{{< tabs title="Position:" style="transparent" color="#131a22ff" >}}
{{% tab title="Start" %}}
{{< code src="assets/examples/lay_anc_ver_pos_sta.txt" lang="go" id="pos" >}}
![center](examples/lay_anc_ver_pos_sta.png)
{{% /tab %}}
{{% tab title="Center" %}}
{{< code src="assets/examples/lay_anc_ver_pos_cen.txt" lang="go" id="pos" >}}
![center](examples/lay_anc_ver_pos_cen.png)
{{% /tab %}}
{{% tab title="End" %}}
{{< code src="assets/examples/lay_anc_ver_pos_end.txt" lang="go" id="pos" >}}
![center](examples/lay_anc_ver_pos_end.png)
{{% /tab %}}
{{< /tabs >}}
