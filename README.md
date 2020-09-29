因为有时要直接下载 github 文件,
内容少的复制一下就行，但很多时候想直接下载下来，然而 raw.\*地址封得太严重，所以要转成 jsdelivr 的 CDN.
反正就是觉得麻烦

就将下载地址转成 cdn 地址再调用 wget，十分简陋,个人平时用用

`./ghdl https://github.com/:owner/:repo/blob/master/:filepath`
