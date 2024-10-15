+++
title = "Row layout"
date = 2024-10-04T19:45:29+03:00
weight = 2
+++

Row layout places all child containers in one row or column. It can be useful for creating lists of widgets.

![preview](examples/lay_row_pre.png)

{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_pre.txt" lang="go" id="full" >}}
{{% /expand %}}

### Layout options

###### Direction

Responsible for whether child containers will follow each other in rows or columns.

{{< tabs title="Direction:" style="transparent" color="#131a22ff" >}}
{{% tab title="Vertical" %}}
{{< code src="assets/examples/lay_row_dir_ver.txt" lang="go" id="dir" >}}
![vert](examples/lay_row_dir_ver.png)
{{% /tab %}}
{{% tab title="Horizontal" %}}
{{< code src="assets/examples/lay_row_dir_hor.txt" lang="go" id="dir" >}}
![horz](examples/lay_row_dir_hor.png)
{{% /tab %}}
{{< /tabs >}}

###### Padding

Layout allows you to specify padding for all child elements but not the itself. Please note that its not possible to specify padding at the *end*, so the *direction* will changed accordingly.

{{< tabs title="Padding:" style="transparent" color="#131a22ff" >}}
{{% tab title="Left" %}}
{{< code src="assets/examples/lay_row_pad_lef.txt" lang="go" id="pad" >}}
![lef](examples/lay_row_pad_lef.png)
{{% /tab %}}
{{% tab title="Right" %}}
{{< code src="assets/examples/lay_row_pad_rig.txt" lang="go" id="pad" >}}
![rig](examples/lay_row_pad_rig.png)
{{% /tab %}}
{{% tab title="Top" %}}
{{< code src="assets/examples/lay_row_pad_top.txt" lang="go" id="pad" >}}
![top](examples/lay_row_pad_top.png)
{{% /tab %}}
{{% tab title="Bottom" %}}
{{< code src="assets/examples/lay_row_pad_bot.txt" lang="go" id="pad" >}}
![top](examples/lay_row_pad_bot.png)
{{% /tab %}}
{{< /tabs >}}

###### Spacing

Layout allows you to specify padding for all child elements but not the itself.

{{< tabs title="Spacing:" style="transparent" color="#131a22ff" >}}
{{% tab title="0" %}}
{{< code src="assets/examples/lay_row_spa_0.txt" lang="go" id="spa" >}}
![0](examples/lay_row_spa_0.png)
{{% /tab %}}
{{% tab title="25" %}}
{{< code src="assets/examples/lay_row_spa_25.txt" lang="go" id="spa" >}}
![25](examples/lay_row_spa_25.png)
{{% /tab %}}
{{% tab title="50" %}}
{{< code src="assets/examples/lay_row_spa_50.txt" lang="go" id="spa" >}}
![50](examples/lay_row_spa_50.png)
{{% /tab %}}
{{% tab title="75" %}}
{{< code src="assets/examples/lay_row_spa_75.txt" lang="go" id="spa" >}}
![75](examples/lay_row_spa_75.png)
{{% /tab %}}
{{< /tabs >}}

### Layout data

###### Stretch

Responsible for stretching the element along the entire length of the opposite axis.

{{< tabs title="Stretch:" style="transparent" color="#131a22ff" >}}
{{% tab title="false" %}}
{{< code src="assets/examples/lay_row_str_fal.txt" lang="go" id="str" >}}
![fal](examples/lay_row_str_fal.png)
{{% /tab %}}
{{% tab title="true" %}}
{{< code src="assets/examples/lay_row_str_tru.txt" lang="go" id="str" >}}
![tru](examples/lay_row_str_tru.png)
{{% /tab %}}
{{< /tabs >}}

###### Position

Responsible for aligning the element along the opposite axis if it is not stretched.

{{< tabs title="Position:" style="transparent" color="#131a22ff" >}}
{{% tab title="Start" %}}
{{< code src="assets/examples/lay_row_pos_sta.txt" lang="go" id="pos" >}}
![sta](examples/lay_row_pos_sta.png)
{{% /tab %}}
{{% tab title="Center" %}}
{{< code src="assets/examples/lay_row_pos_cen.txt" lang="go" id="pos" >}}
![sta](examples/lay_row_pos_cen.png)
{{% /tab %}}
{{% tab title="End" %}}
{{< code src="assets/examples/lay_row_pos_end.txt" lang="go" id="pos" >}}
![sta](examples/lay_row_pos_end.png)
{{% /tab %}}
{{< /tabs >}}


###### Max size

Responsible for the allowable size of the container.

{{< tabs title="Max:" style="transparent" color="#131a22ff" >}}
{{% tab title="Width" %}}
{{< code src="assets/examples/lay_row_max_wid.txt" lang="go" id="max" >}}
![wid](examples/lay_row_max_wid.png)
{{% /tab %}}
{{% tab title="Height" %}}
{{< code src="assets/examples/lay_row_max_hei.txt" lang="go" id="max" >}}
![hei](examples/lay_row_max_hei.png)
{{% /tab %}}
{{< /tabs >}}
