<div align="center" style="text-align: center;">
    <h1>gitlab-go</h1>
    <h3>GitLab Go 脚手架</h3>
    <a target="_blank" href="https://github.com/996icu/996.ICU/blob/master/LICENSE">
        <img alt="License-Anti" src="https://img.shields.io/badge/License-Anti 996-blue.svg">
    </a>
    <a target="_blank" href="https://996.icu/#/zh_CN">
        <img alt="Link-996" src="https://img.shields.io/badge/Link-996.icu-red.svg">
    </a>
    <a target="_blank" href="https://qm.qq.com/cgi-bin/qm/qr?k=ZieC6s1WB4njfVbrDHYgoNS8YpT26VtF&jump_from=webapi">
        <img alt="QQ群" src="https://img.shields.io/badge/QQ群-696503132-blue.svg"/>
    </a>
</div>

<p></p>

<div align="center" style="text-align: center;">
    <a target="_blank" href="https://work.weixin.qq.com/gm/75cfc47d6a341047e4b6aca7389bdfa8">
        <img alt="企业微信群" src="static/wechat-work.jpg" height="100"/>
    </a>
</div>

<p></p>

<div align="center" style="text-align: center;">
    基于 go 语言的 gitlab 脚手架
</div>

<p></p>

<div align="center" style="text-align: center;">
  为简化开发工作、提高生产率、解决常见问题而生
</div>

<p></p>

<div align="center" style="text-align: center;">

  <a target="_blank" href="https://space.bilibili.com/198580655">
    <img alt="bilibili 粉丝" src="https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fapi.spencerwoo.com%2Fsubstats%2F%3Fsource%3Dbilibili%26queryKey%3D198580655&label=bilibili%20fans&query=%24.data.totalSubs&logo=bilibili">
  </a>

  <a target="_blank" href="https://blog.csdn.net/qq_32596527">
    <img alt="CSDN 码龄" src="https://img.shields.io/badge/dynamic/xml?color=orange&label=CSDN&query=%2F%2Fdiv%5B%40class%3D%27person-code-age%27%5D%5B1%5D%2Fspan%5B1%5D%2Ftext%28%29%5B1%5D&url=https%3A%2F%2Fblog.csdn.net%2Fqq_32596527&logo=data:image/x-icon;base64,AAABAAEAICAAAAEAIACoEAAAFgAAACgAAAAgAAAAQAAAAAEAIAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAxVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zda/P9qhPz/mKr9/7bC/f/Fz/7/ydL+/8HM/v+tu/3/jaH9/156/P8zV/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/z9h/P+gsP3/8fP+/////////////////////////////////////////////////+ru/v+Zqv3/PV/8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9lgPz/6+/+///////////////////////////////////////////////////////////////////////s7/7/Y378/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aoT8//r6/v///////////////////////v7+/+Po/v/R2f7/y9T+/9rg/v/3+f7////////////////////////////j6P7/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Zm/P/w8/7/////////////////5+v+/4ab/f9AYvz/MVX8/zFV/P8xVfz/MVX8/zVY/P9kf/z/tsP9//39/v////////////T2/v8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/sL79/////////////////87W/v8/Yfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ZYD8//L0/v//////n7D9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Bh/P/6+/7////////////v8v7/QmP8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/TWz8/3GJ/P8yVvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/e5L8/////////////////5qr/f8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P+mtv3/////////////////XHn8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7/L/f////////////////87Xfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ydL+////////////+/v+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P/Ezv7////////////9/f7/M1b8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7G//f////////////////9HZ/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kqX9/////////////////22H/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9kf/z/////////////////pbX9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zRX/P/v8v7////////////s7/7/Nln8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/6Ky/f////////////////+Inf3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/RWb8//f4/v////////////H0/v9Kafz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/PV/8/1Jw/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kKT9/////////////////9vh/v9DZPz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1Fv/P/m6/7//v7+/3aO/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8zVvz/xM79/////////////////+fr/v9viPz/MVX8/zFV/P8xVfz/MVX8/zRX/P+Emf3/8/X+////////////xc/+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P87Xfz/ztf+///////////////////////i5/7/sL79/5+w/f+ywP3/6u3+//////////////////////+uvP3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P83Wvz/sL79//7+/v//////////////////////////////////////////////////////3OL+/0Vl/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aYP8/9Pb/v//////////////////////////////////////9fb+/5yu/f84W/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1d0/P+Spf3/t8T9/8fR/v/Dzv7/qrn9/3uS/P88Xvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=">
  </a>

  <a target="_blank" href="https://blog.csdn.net/qq_32596527">
    <img alt="CSDN 粉丝" src="https://img.shields.io/badge/dynamic/xml?color=orange&label=CSDN&prefix=%E7%B2%89%E4%B8%9D&query=%2F%2Fli%5B4%5D%2Fa%5B1%5D%2Fdiv%5B%40class%3D%27user-profile-statistics-num%27%5D%5B1%5D%2Ftext%28%29%5B1%5D&url=https%3A%2F%2Fblog.csdn.net%2Fqq_32596527&logo=data:image/x-icon;base64,AAABAAEAICAAAAEAIACoEAAAFgAAACgAAAAgAAAAQAAAAAEAIAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAxVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zda/P9qhPz/mKr9/7bC/f/Fz/7/ydL+/8HM/v+tu/3/jaH9/156/P8zV/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/z9h/P+gsP3/8fP+/////////////////////////////////////////////////+ru/v+Zqv3/PV/8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9lgPz/6+/+///////////////////////////////////////////////////////////////////////s7/7/Y378/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aoT8//r6/v///////////////////////v7+/+Po/v/R2f7/y9T+/9rg/v/3+f7////////////////////////////j6P7/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Zm/P/w8/7/////////////////5+v+/4ab/f9AYvz/MVX8/zFV/P8xVfz/MVX8/zVY/P9kf/z/tsP9//39/v////////////T2/v8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/sL79/////////////////87W/v8/Yfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ZYD8//L0/v//////n7D9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Bh/P/6+/7////////////v8v7/QmP8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/TWz8/3GJ/P8yVvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/e5L8/////////////////5qr/f8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P+mtv3/////////////////XHn8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7/L/f////////////////87Xfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ydL+////////////+/v+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P/Ezv7////////////9/f7/M1b8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7G//f////////////////9HZ/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kqX9/////////////////22H/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9kf/z/////////////////pbX9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zRX/P/v8v7////////////s7/7/Nln8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/6Ky/f////////////////+Inf3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/RWb8//f4/v////////////H0/v9Kafz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/PV/8/1Jw/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kKT9/////////////////9vh/v9DZPz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1Fv/P/m6/7//v7+/3aO/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8zVvz/xM79/////////////////+fr/v9viPz/MVX8/zFV/P8xVfz/MVX8/zRX/P+Emf3/8/X+////////////xc/+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P87Xfz/ztf+///////////////////////i5/7/sL79/5+w/f+ywP3/6u3+//////////////////////+uvP3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P83Wvz/sL79//7+/v//////////////////////////////////////////////////////3OL+/0Vl/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aYP8/9Pb/v//////////////////////////////////////9fb+/5yu/f84W/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1d0/P+Spf3/t8T9/8fR/v/Dzv7/qrn9/3uS/P88Xvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=">
  </a>

  <a target="_blank" href="https://blog.csdn.net/qq_32596527">
    <img alt="CSDN 访问" src="https://img.shields.io/badge/dynamic/xml?color=orange&label=CSDN&prefix=%E8%AE%BF%E9%97%AE&query=//span[1]/div[@class='user-profile-statistics-num'][1]/text()[1]&url=https%3A%2F%2Fblog.csdn.net%2Fqq_32596527&logo=data:image/x-icon;base64,AAABAAEAICAAAAEAIACoEAAAFgAAACgAAAAgAAAAQAAAAAEAIAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAxVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zda/P9qhPz/mKr9/7bC/f/Fz/7/ydL+/8HM/v+tu/3/jaH9/156/P8zV/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/z9h/P+gsP3/8fP+/////////////////////////////////////////////////+ru/v+Zqv3/PV/8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9lgPz/6+/+///////////////////////////////////////////////////////////////////////s7/7/Y378/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aoT8//r6/v///////////////////////v7+/+Po/v/R2f7/y9T+/9rg/v/3+f7////////////////////////////j6P7/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Zm/P/w8/7/////////////////5+v+/4ab/f9AYvz/MVX8/zFV/P8xVfz/MVX8/zVY/P9kf/z/tsP9//39/v////////////T2/v8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/sL79/////////////////87W/v8/Yfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ZYD8//L0/v//////n7D9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Bh/P/6+/7////////////v8v7/QmP8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/TWz8/3GJ/P8yVvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/e5L8/////////////////5qr/f8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P+mtv3/////////////////XHn8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7/L/f////////////////87Xfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ydL+////////////+/v+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P/Ezv7////////////9/f7/M1b8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7G//f////////////////9HZ/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kqX9/////////////////22H/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9kf/z/////////////////pbX9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zRX/P/v8v7////////////s7/7/Nln8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/6Ky/f////////////////+Inf3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/RWb8//f4/v////////////H0/v9Kafz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/PV/8/1Jw/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kKT9/////////////////9vh/v9DZPz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1Fv/P/m6/7//v7+/3aO/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8zVvz/xM79/////////////////+fr/v9viPz/MVX8/zFV/P8xVfz/MVX8/zRX/P+Emf3/8/X+////////////xc/+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P87Xfz/ztf+///////////////////////i5/7/sL79/5+w/f+ywP3/6u3+//////////////////////+uvP3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P83Wvz/sL79//7+/v//////////////////////////////////////////////////////3OL+/0Vl/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aYP8/9Pb/v//////////////////////////////////////9fb+/5yu/f84W/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1d0/P+Spf3/t8T9/8fR/v/Dzv7/qrn9/3uS/P88Xvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=">
  </a>

  <a target="_blank" href="https://blog.csdn.net/qq_32596527">
    <img alt="CSDN 博客" src="https://img.shields.io/badge/dynamic/json?color=orange&label=CSDN&prefix=%E5%8D%9A%E5%AE%A2&query=%24.data.blog&suffix=%E7%AF%87&url=https%3A%2F%2Fblog.csdn.net%2Fcommunity%2Fhome-api%2Fv1%2Fget-tab-total%3Fusername%3Dqq_32596527&logo=data:image/x-icon;base64,AAABAAEAICAAAAEAIACoEAAAFgAAACgAAAAgAAAAQAAAAAEAIAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAxVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zda/P9qhPz/mKr9/7bC/f/Fz/7/ydL+/8HM/v+tu/3/jaH9/156/P8zV/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/z9h/P+gsP3/8fP+/////////////////////////////////////////////////+ru/v+Zqv3/PV/8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9lgPz/6+/+///////////////////////////////////////////////////////////////////////s7/7/Y378/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aoT8//r6/v///////////////////////v7+/+Po/v/R2f7/y9T+/9rg/v/3+f7////////////////////////////j6P7/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Zm/P/w8/7/////////////////5+v+/4ab/f9AYvz/MVX8/zFV/P8xVfz/MVX8/zVY/P9kf/z/tsP9//39/v////////////T2/v8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/sL79/////////////////87W/v8/Yfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ZYD8//L0/v//////n7D9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Bh/P/6+/7////////////v8v7/QmP8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/TWz8/3GJ/P8yVvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/e5L8/////////////////5qr/f8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P+mtv3/////////////////XHn8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7/L/f////////////////87Xfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ydL+////////////+/v+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P/Ezv7////////////9/f7/M1b8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7G//f////////////////9HZ/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kqX9/////////////////22H/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9kf/z/////////////////pbX9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zRX/P/v8v7////////////s7/7/Nln8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/6Ky/f////////////////+Inf3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/RWb8//f4/v////////////H0/v9Kafz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/PV/8/1Jw/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kKT9/////////////////9vh/v9DZPz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1Fv/P/m6/7//v7+/3aO/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8zVvz/xM79/////////////////+fr/v9viPz/MVX8/zFV/P8xVfz/MVX8/zRX/P+Emf3/8/X+////////////xc/+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P87Xfz/ztf+///////////////////////i5/7/sL79/5+w/f+ywP3/6u3+//////////////////////+uvP3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P83Wvz/sL79//7+/v//////////////////////////////////////////////////////3OL+/0Vl/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aYP8/9Pb/v//////////////////////////////////////9fb+/5yu/f84W/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1d0/P+Spf3/t8T9/8fR/v/Dzv7/qrn9/3uS/P88Xvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=">
  </a>

  <a target="_blank" href="https://www.jetbrains.com/idea">
    <img alt="IntelliJ IDEA" src="https://img.shields.io/static/v1?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAAAAAByaaZbAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAHdElNRQfmBRkBICRBfW8eAAABPklEQVRIx+2UTStEYRiGr/kqZWaKMoyQUsoCWdn4AZKPlZSF5fgPVhYWStlZqPkB7GynWSg/QCmhUFOUGCPNZBrlnNtijmneOYZzykrn2p2e536f577P2wsBPvmQlACQg13Mr4b9CCTpuMunQIdGT8gQRCBZAQRPu0B6aRjsvqKXCRcAJO8lTTf3/GQJKF8Bz5481CfMvEnnxtrRduKhPET7R6GWka+UrBW//8Hej3haqZQFYgNz8VDmYbNNT8iSFDdMM1iSyh3uWHsAesNgVc1D7k4hMeISrN0uAOtAwTYF6Smg2UQUYDYbO8qdjS0Cua9CahsgPd8NleuW3WOFRiTV+nTj8mnL5XbixinVlnEJ7L1vkuzcuJT0ejDufL98UTjZmWyJsqFJvT9aBAT8J5zrrXYFF788xn8gCPDCJ2cr3I1zqSjOAAAAJXRFWHRkYXRlOmNyZWF0ZQAyMDIyLTA1LTI1VDAxOjMyOjM2KzAwOjAwH/0yeQAAACV0RVh0ZGF0ZTptb2RpZnkAMjAyMi0wNS0yNVQwMTozMjozNiswMDowMG6gisUAAAAASUVORK5CYII=&message=IntelliJ IDEA">
  </a>

  <a target="_blank" href="https://github.com/xuxiaowei-com-cn/gitlab-go">
    <img alt="GitHub stars" src="https://img.shields.io/github/stars/xuxiaowei-com-cn/gitlab-go?logo=github">
  </a>

  <a target="_blank" href="https://github.com/xuxiaowei-com-cn/gitlab-go">
    <img alt="GitHub forks" src="https://img.shields.io/github/forks/xuxiaowei-com-cn/gitlab-go?logo=github">
  </a>

  <a target="_blank" href="https://github.com/xuxiaowei-com-cn/gitlab-go">
    <img alt="GitHub watchers" src="https://img.shields.io/github/watchers/xuxiaowei-com-cn/gitlab-go?logo=github">
  </a>

  <a target="_blank" href="https://github.com/xuxiaowei-com-cn/gitlab-go">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/xuxiaowei-com-cn/gitlab-go">
  </a>

  <a target="_blank" href="https://gitee.com/xuxiaowei-com-cn/gitlab-go">
    <img alt="码云Gitee stars" src="https://gitee.com/xuxiaowei-com-cn/gitlab-go/badge/star.svg?theme=blue">
  </a>

  <a target="_blank" href="https://gitee.com/xuxiaowei-com-cn/gitlab-go">
    <img alt="码云Gitee forks" src="https://gitee.com/xuxiaowei-com-cn/gitlab-go/badge/fork.svg?theme=blue">
  </a>

  <a target="_blank" href="https://gitlab.com/xuxiaowei-com-cn/gitlab-go">
    <img alt="Gitlab stars" src="https://badgen.net/gitlab/stars/xuxiaowei-com-cn/gitlab-go?icon=gitlab">
  </a>

  <a target="_blank" href="https://gitlab.com/xuxiaowei-com-cn/gitlab-go">
    <img alt="Gitlab forks" src="https://badgen.net/gitlab/forks/xuxiaowei-com-cn/gitlab-go?icon=gitlab">
  </a>

  <a target="_blank" href="https://github.com/xuxiaowei-com-cn/gitlab-go">
    <img alt="OSCS Status" src="https://www.oscs1024.com/platform/badge/xuxiaowei-com-cn/gitlab-go.svg?size=small">
  </a>

  <a target="_blank" href="https://github.com/xuxiaowei-com-cn/gitlab-go">
    <img alt="total lines" src="https://tokei.rs/b1/github/xuxiaowei-com-cn/gitlab-go">
  </a>

  <a target="_blank" href="https://www.apache.org/licenses/LICENSE-2.0">
    <img alt="code style" src="https://img.shields.io/badge/license-Apache 2-blue">
  </a>
</div>

## 支持的系统

1. Linux amd64、arm64
2. Windows amd64、arm64
3. Mac amd64、arm64
4. LoongArch 64-bit

# 构建

本项目发布在 GitHub 进行构建，并使用构建后的程序将自己发布到各大代码托管平台，发布过程参见：
[GitHub Actions](https://github.com/xuxiaowei-com-cn/gitlab-go/actions/workflows/go-push.yml)

## 下载

1. [gitee](https://gitee.com/xuxiaowei-com-cn/gitlab-go/releases)
2. [gitlab](https://gitlab.com/xuxiaowei-com-cn/gitlab-go/-/releases)
3. [github](https://github.com/xuxiaowei-com-cn/gitlab-go/releases)
4. [gitlink](https://gitlink.org.cn/xuxiaowei-com-cn/gitlab-go/releases)

## 支持的功能

1. 一键发布到 github，可包含产物上传
2. 一键发布到 gitlab，可包含产物上传，可自定义域名（支持自建 gitlab），支持将产物文件名、链接导出为 map，可供 gitee 使用
3. 一键发布到 gitee，由于 gitee 暂不支持提供上传产物的 API 接口，
   本工具支持提供 json 文件（map 形式，键：代表文件名，值：代表下载链接）作为产物，
   本项目使用 [GitLink](https://www.gitlink.org.cn) 作为 gitee 产物链接
4. 一键发布到 gitlink，可包含产物上传（需要等到官方开放 token 功能，或者联系官方人员申请 token 才能使用），
   本工具支持提供 json 文件（map 形式，键：代表文件名，值：代表下载链接）作为产物

## 开发命令

### get

```shell
go env -w GOPROXY=https://goproxy.cn,direct
# go env -w GOPROXY=https://proxy.golang.org,direct
# go env -w GOPROXY=https://goproxy.io,direct
# go env -w GOPROXY=https://mirrors.aliyun.com/goproxy,direct
# go env -w GOPROXY=https://mirrors.cloud.tencent.com/go,direct
go get -u github.com/urfave/cli/v2
go get -u github.com/xanzy/go-gitlab
go get -u github.com/xuxiaowei-com-cn/git-go@main
go get -u gopkg.in/yaml.v3
```

### mod

```shell
go mod tidy
```

```shell
go mod download
```

### run

```shell
go run main.go
```

### run help

- Windows 环境为 %xxx%
- Linux 环境为 $xxx

```shell
go run main.go --help
```

```shell
$ go run main.go --help
NAME:
   gitlab-go - 基于 Go 语言开发的 GitLab 命令行工具

USAGE:
   gitlab-go [global options] command [command options]

VERSION:
   dev

AUTHOR:
   徐晓伟 <xuxiaowei@xuxiaowei.com.cn>

COMMANDS:
   access-request, access-requests, ar                              群组和项目访问请求 API，中文文档：https://docs.gitlab.cn/jh/api/access_requests.html
   board, boards                                                    项目议题板 API，中文文档：https://docs.gitlab.cn/jh/api/boards.html
   container-registry, cr                                           容器仓库 API，中文文档：https://docs.gitlab.cn/jh/api/container_registry.html
   environments, environment, env                                   环境 API，中文文档：https://docs.gitlab.cn/jh/api/environments.html
   instance-level-ci-variables, instance-level-ci-variable, ilcv    实例级 CI/CD 变量 API，中文文档：https://docs.gitlab.cn/jh/api/instance_level_ci_variables.html
   issue, issues                                                    议题 API，中文文档：https://docs.gitlab.cn/jh/api/issues.html
   job-artifact, job-artifacts, ja                                  作业产物 API，中文文档：https://docs.gitlab.cn/jh/api/job_artifacts.html
   job, jobs, j                                                     作业 API，中文文档：https://docs.gitlab.cn/jh/api/jobs.html
   pipeline, pipelines, pl                                          流水线 API，中文文档：https://docs.gitlab.cn/jh/api/pipelines.html
   project-level-variables, project-level-variable, plv             项目级别 CI/CD 变量 API，中文文档：https://docs.gitlab.cn/jh/api/project_level_variables.html
   project, projects, p                                             项目 API，中文文档：https://docs.gitlab.cn/jh/api/projects.html
   mix-archive                                                      归档（混合命令，多接口命令）
   mix-delete, mix-rm                                               删除（混合命令，多接口命令）
   mix-create-environments, mix-create-environment, mix-create-env  创建新环境（混合命令，多接口命令）
   mix-export                                                       导出（混合命令，多接口命令）
   mix-transfer                                                     转移（混合命令，多接口命令）
   mix-unarchive                                                    取消归档（混合命令，多接口命令）
   help, h                                                          Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

COPYRIGHT:
   徐晓伟工作室 <xuxiaowei@xuxiaowei.com.cn>
```

- [access-request - 群组和项目访问请求 API](https://docs.gitlab.cn/jh/api/access_requests.html)

    ```shell
    $ go run main.go access-request --help
    NAME:
       gitlab-go access-request - 群组和项目访问请求 API，中文文档：https://docs.gitlab.cn/jh/api/access_requests.html
    
    USAGE:
       gitlab-go access-request command [command options]
    
    COMMANDS:
       group, groups      为群组列出访问请求
       project, projects  为项目列出访问请求
       help, h            Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value  实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value     your_access_token
       --page value      页码（默认：1），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 1)
       --per-page value  每页列出的项目数（默认：20；最大：100），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 20)
       --print-json      打印 JSON (default: false)
       --print-time      打印时间 (default: false)
       --id value        项目 ID 或 URL 编码的路径
       --help, -h        show help
    ```

- [board - 项目议题板 API](https://docs.gitlab.cn/jh/api/boards.html)

    ```shell
    $ go run main.go board --help
    NAME:
       gitlab-go board - 项目议题板 API，中文文档：https://docs.gitlab.cn/jh/api/boards.html
    
    USAGE:
       gitlab-go board command [command options]
    
    COMMANDS:
       list     列出项目议题板
       help, h  Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value  实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value     your_access_token
       --page value      页码（默认：1），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 1)
       --per-page value  每页列出的项目数（默认：20；最大：100），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 20)
       --print-json      打印 JSON (default: false)
       --print-time      打印时间 (default: false)
       --id value        项目 ID 或 URL 编码的路径
       --help, -h        show help
    ```

- [container-registry - 容器仓库 API](https://docs.gitlab.cn/jh/api/container_registry.html)

    ```shell
    $ go run main.go container-registry --help
    NAME:
       gitlab-go container-registry - 容器仓库 API，中文文档：https://docs.gitlab.cn/jh/api/container_registry.html
    
    USAGE:
       gitlab-go container-registry command [command options]
    
    COMMANDS:
       list                列出仓库内存储库
       list-tags           列出仓库里存储库的标签
       get-tag             获取仓库里存储库的某个标签的详情
       delete-tag, rm-tag  删除仓库里存储库的某个标签
       help, h             Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value    实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value       your_access_token
       --page value        页码（默认：1），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 1)
       --per-page value    每页列出的项目数（默认：20；最大：100），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 20)
       --print-json        打印 JSON (default: false)
       --print-time        打印时间 (default: false)
       --id value          项目 ID 或 URL 编码的路径
       --repository value  仓库里存储库的 ID
       --tag-name value    标签的名称
       --help, -h          show help
    ```

- [environments 环境 API](https://docs.gitlab.cn/jh/api/environments.html)

    ```shell
    $ go run main.go environments --help
    NAME:
       gitlab-go environments - 环境 API，中文文档：https://docs.gitlab.cn/jh/api/environments.html
    
    USAGE:
       gitlab-go environments command [command options]
    
    COMMANDS:
       list     列举环境
       create   创建新环境
       help, h  Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value      实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value         your_access_token
       --id value            项目 ID 或 URL 编码的路径
       --name value          环境名称
       --external_url value  该环境的链接位置
       --tier value          新环境的层级。允许设置的值为 production， staging， testing， development 和 other
       --page value          页码（默认：1），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 1)
       --per-page value      每页列出的项目数（默认：20；最大：100），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 20)
       --print-json          打印 JSON (default: false)
       --print-time          打印时间 (default: false)
       --help, -h            show help
    ```

- [instance-level-ci-variables - 实例级 CI/CD 变量 API](https://docs.gitlab.cn/jh/api/instance_level_ci_variables.html)

    ```shell
    $ go run main.go instance-level-ci-variables --help
    NAME:
       gitlab-go instance-level-ci-variables - 实例级 CI/CD 变量 API，中文文档：https://docs.gitlab.cn/jh/api/instance_level_ci_variables.html
    
    USAGE:
       gitlab-go instance-level-ci-variables command [command options]
    
    COMMANDS:
       list     列出所有实例变量
       help, h  Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value  实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value     your_access_token
       --page value      页码（默认：1），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 1)
       --per-page value  每页列出的项目数（默认：20；最大：100），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 20)
       --print-json      打印 JSON (default: false)
       --print-time      打印时间 (default: false)
       --help, -h        show help
    ```

- [issue - 议题 API](https://docs.gitlab.cn/jh/api/issues.html)

    ```shell
    $ go run main.go issue --help
    NAME:
       gitlab-go issue - 议题 API，中文文档：https://docs.gitlab.cn/jh/api/issues.html
    
    USAGE:
       gitlab-go issue command [command options]
    
    COMMANDS:
       list                    列出议题
       delete, rm              删除议题
       delete-range, rm-range  删除议题（范围）
       help, h                 Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value                                   实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value                                      your_access_token
       --page value                                       页码（默认：1），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 1)
       --per-page value                                   每页列出的项目数（默认：20；最大：100），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 20)
       --print-json                                       打印 JSON (default: false)
       --print-time                                       打印时间 (default: false)
       --recursion                                        递归 (default: false)
       --assignee-username value                          对于给定的用户名 username，返回指派给这个用户的议题。与 assignee_id 相似且与其冲突。在免费版中，assignee_username 数组只能包含单个成员，否则将报参数错误提示。
       --author-id value                                  对于给定的用户 id，返回这个用户创建的议题。与 author_username 冲突。与 scope=all 或 scope=assigned_to_me 配合使用。
       --author-username value                            对于给定的用户名 username，返回这个用户创建的议题。与 author_id 相似且与其冲突。
       --confidential                                     筛选私密议题与公开议题。 (default: false)
       --due-date value                                   返回没有截止日期、已经逾期、本周内逾期、本月内逾期或介于两周前和下个月之间逾期的议题。可接受的值：0（没有截止日期）、any、today、tomorrow、overdue、week、month、next_month_and_previous_two_weeks。
       --iids value                                       返回包含给定 iid 的议题。
       --in value                                         修改 search 属性的范围。可以使用 title、description 或使用半角逗号对他们进行连接。默认值是 title,description。 (default: "title,description")
       --issue-type value                                 筛选议题的类型，可选值为 issue、incident 或 test_case。引入于 13.12 版本。
       --iteration-id value                               对于给定的迭代 ID，返回与这个迭代关联的议题。使用 None 则返回未与迭代关联的议题。使用 Any 则返回存在关联迭代的议题。引入于 13.6 版本。
       --milestone value                                  里程碑名称。使用 None 则列出没有里程碑的议题。使用 Any 则列出存在关联里程碑的议题。None 及 Any 的用法将会在未来被弃用，请使用 milestone_id 替代。milestone 与 milestone_id 冲突。
       --milestone-id value                               对于给定的时间段（None、Any、Upcoming 或 Started），返回与该时间段里程碑相关联的议题。使用 None 则列出没有里程碑的议题。使用 Any 则列出存在关联里程碑的议题。使用 Upcoming 则列出与未开始里程碑相关联的议题。使用 Started 则列出与已开始里程碑相关联的议题。milestone 和 milestone_id 冲突。引入于 14.3 版本。
       --my-reaction-emoji value                          对于给定的 emoji，返回用户使用该表情回应的议题。使用 None 则返回没有使用表情回应的议题。使用 Any 则返回使用至少一个表情回应的议题。
       --order-by value                                   返回根据 created_at、due_date、label_priority、milestone_due、popularity、priority、relative_position、title、updated_at 或 weight 排序的议题。默认值是 created_at。 (default: "created_at")
       --scope value                                      返回满足范围 created_by_me、assigned_to_me 或 all 的议题。默认值是 created_by_me。 (default: "created_by_me")
       --search value                                     根据 title 和 description 搜索议题。
       --sort value                                       按照 asc 或者 desc 排序 (default: "desc")
       --state value                                      返回全部 all 议题或仅返回处于 opened 或 closed 状态的议题。 (default: "all")
       --with-labels-details                              若为 true 则返回更详尽的标签信息：:name、:color、:description、:description_html、:text_color。默认值是 false。description_html 属性引入于 12.7 版本。 (default: false)
       --created-after value                              对于给定的时间戳，返回不早于该时间创建的议题。时间戳应符合 ISO 8601 格式（2019-03-15T08:00:00Z）
       --created-before value                             对于给定的时间戳，返回不晚于该时间创建的议题。时间戳应符合 ISO 8601 格式（2019-03-15T08:00:00Z）。
       --issue-id value                                   项目议题的内部 ID
       --issue-id-range value [ --issue-id-range value ]  议题ID的范围，支持范围如下：
                                                          单数：1
                                                          多个数字（使用英文逗号隔开）：1,2,3,7,8,15
                                                          支持范围：5-10,
                                                          支持范围方向选择：-10（小于等于10，即：从 0 到 10），214-（大于等于214，即：从 214 到 214 + 10000，数据范围不超过 10000）
       --help, -h                                         show help
    ```

- [job-artifact - 作业产物 API](https://docs.gitlab.cn/jh/api/job_artifacts.html)

    ```shell
    $ go run main.go job-artifact --help
    NAME:
       gitlab-go job-artifact - 作业产物 API，中文文档：https://docs.gitlab.cn/jh/api/job_artifacts.html
    
    USAGE:
       gitlab-go job-artifact command [command options]
    
    COMMANDS:
       get                                    获取（下载）作业产物
       download, dl                           下载产物归档文件（未完成）
       delete, rm                             删除作业产物
       delete-project, delete-projects, rm-p  删除项目产物（计划删除，如需立即删除请使用混合命令）
       help, h                                Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value        实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value           your_access_token
       --id value              项目 ID 或 URL 编码的路径
       --job-id value          作业 ID
       --artifacts-name value  保存产物名称（保存到系统磁盘的名称） (default: "artifacts.zip")
       --help, -h              show help
    ```

- [job - 作业 API](https://docs.gitlab.cn/jh/api/jobs.html)

    ```shell
    $ go run main.go job --help
    NAME:
       gitlab-go job - 作业 API，中文文档：https://docs.gitlab.cn/jh/api/jobs.html
    
    USAGE:
       gitlab-go job command [command options]
    
    COMMANDS:
       list     列出项目作业
       erase    删除作业（删除作业产物和作业日志）
       help, h  Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value                               实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value                                  your_access_token
       --page value                                   页码（默认：1），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 1)
       --per-page value                               每页列出的项目数（默认：20；最大：100），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 20)
       --print-json                                   打印 JSON (default: false)
       --print-time                                   打印时间 (default: false)
       --recursion                                    递归 (default: false)
       --sort value                                   按照 asc 或者 desc 排序 (default: "desc")
       --id value                                     项目 ID 或 URL 编码的路径
       --job-id-range value [ --job-id-range value ]  Job ID的范围，支持范围如下：
                                                      单数：1
                                                      多个数字（使用英文逗号隔开）：1,2,3,7,8,15
                                                      支持范围：5-10,
                                                      支持范围方向选择：-10（小于等于10，即：从 0 到 10），214-（大于等于214，即：从 214 到 214 + 10000，数据范围不超过 10000）
       --scope value                                  要显示的作业范围。以下之一或数组：created、pending、running、failed、success、canceled、skipped、waiting_for_resource 或 manual。范围如果未提供，则返回所有作业。
       --help, -h                                     show help
    ```

- [pipeline - 流水线 API](https://docs.gitlab.cn/jh/api/pipelines.html)

    ```shell
    $ go run main.go pipeline --help
    NAME:
       gitlab-go pipeline - 流水线 API，中文文档：https://docs.gitlab.cn/jh/api/pipelines.html
    
    USAGE:
       gitlab-go pipeline command [command options]
    
    COMMANDS:
       list     列出项目流水线
       help, h  Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value  实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value     your_access_token
       --sort value      按照 asc 或者 desc 排序 (default: "desc")
       --page value      页码（默认：1），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 1)
       --per-page value  每页列出的项目数（默认：20；最大：100），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 20)
       --print-json      打印 JSON (default: false)
       --print-time      打印时间 (default: false)
       --recursion       递归 (default: false)
       --id value        项目 ID 或 URL 编码的路径
       --help, -h        show help
    ```

- [project-level-variables 项目级别 CI/CD 变量 API](https://docs.gitlab.cn/jh/api/project_level_variables.html)

    ```shell
    $ go run main.go project-level-variables --help
    NAME:
       gitlab-go project-level-variables - 项目级别 CI/CD 变量 API，中文文档：https://docs.gitlab.cn/jh/api/project_level_variables.html
    
    USAGE:
       gitlab-go project-level-variables command [command options]
    
    COMMANDS:
       list     列出项目变量
       help, h  Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value  实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value     your_access_token
       --page value      页码（默认：1），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 1)
       --per-page value  每页列出的项目数（默认：20；最大：100），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 20)
       --print-json      打印 JSON (default: false)
       --print-time      打印时间 (default: false)
       --help, -h        show help
    ```

- [project - 项目 API](https://docs.gitlab.cn/jh/api/projects.html)

    ```shell
    $ go run main.go project --help
    NAME:
       gitlab-go project - 项目 API，中文文档：https://docs.gitlab.cn/jh/api/projects.html
    
    USAGE:
       gitlab-go project command [command options]
    
    COMMANDS:
       list     列出所有项目
       help, h  Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value     实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value        your_access_token
       --sort value         按照 asc 或者 desc 排序 (default: "desc")
       --page value         页码（默认：1），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 1)
       --per-page value     每页列出的项目数（默认：20；最大：100），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 20)
       --print-json         打印 JSON (default: false)
       --print-time         打印时间 (default: false)
       --search value       根据 title 和 description 搜索议题。
       --search-namespaces  匹配搜索条件时包括上级命名空间。默认为 false。 (default: false)
       --order-by value     返回按 id、name、path、created_at、updated_at、last_activity_at 或 similarity 字段排序的项目。repository_size、storage_size、packages_size 或 wiki_size 字段只允许管理员使用。similarity（引入于 14.1 版本）仅在搜索 时可用，并且仅限于当前用户所属的项目。默认是created_at。 (default: "created_at")
       --help, -h           show help
    ```

- mix-archive 归档（混合命令，多接口命令）

    ```shell
    $ go run main.go mix-archive --help
    NAME:
       gitlab-go mix-archive - 归档（混合命令，多接口命令）
    
    USAGE:
       gitlab-go mix-archive command [command options]
    
    COMMANDS:
       all      归档所有项目（混合命令，多接口命令）
       help, h  Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value  实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value     your_access_token
       --owned           当前用户明确拥有的项目。 (default: false)
       --help, -h        show help
    ```

- mix-delete 删除（混合命令，多接口命令）

    ```shell
    $ go run main.go mix-delete --help
    NAME:
       gitlab-go mix-delete - 删除（混合命令，多接口命令）
    
    USAGE:
       gitlab-go mix-delete command [command options]
    
    COMMANDS:
       artifact, artifacts          根据项目路径/ID、流水线IID范围删除产物（混合命令，多接口命令，立即删除）
       all-artifact, all-artifacts  根据项目路径/ID删除所有产物（混合命令，多接口命令，立即删除）
       job, jobs                    根据项目路径/ID、流水线IID范围删除作业产物和作业日志（混合命令，多接口命令，立即删除）
       all-job, all-jobs            根据项目路径/ID删除所有作业产物和作业日志（混合命令，多接口命令，立即删除）
       help, h                      Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value                         实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value                            your_access_token
       --sort value                             按照 asc 或者 desc 排序 (default: "desc")
       --page value                             页码（默认：1），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 1)
       --per-page value                         每页列出的项目数（默认：20；最大：100），中文文档 https://docs.gitlab.cn/jh/api/rest/index.html#pagination (default: 20)
       --id value                               项目 ID 或 URL 编码的路径
       --iid-range value [ --iid-range value ]  流水线ID的范围，支持范围如下：
                                                单数：1
                                                多个数字（使用英文逗号隔开）：1,2,3,7,8,15
                                                支持范围：5-10,
                                                支持范围方向选择：-10（小于等于10，即：从 0 到 10），214-（大于等于214，即：从 214 到 214 + 10000，数据范围不超过 10000）
       --help, -h                               show help
    ```

- mix-create-environments 创建新环境（混合命令，多接口命令）

    ```shell
    $ go run main.go mix-create-environments --help
    NAME:
       gitlab-go mix-create-environments - 创建新环境（混合命令，多接口命令）
    
    USAGE:
       gitlab-go mix-create-environments command [command options]
    
    COMMANDS:
       all, a   所有项目创建新环境
       help, h  Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value                                                   实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value                                                      your_access_token
       --owned                                                            当前用户明确拥有的项目。 (default: false)
       --skip-project-path value [ --skip-project-path value ]            跳过项目路径
       --skip-project-wiki-path value [ --skip-project-wiki-path value ]  跳过项目wiki路径
       --allow-failure                                                    允许失败 (default: false)
       --name value                                                       环境名称
       --external_url value                                               该环境的链接位置
       --tier value                                                       新环境的层级。允许设置的值为 production， staging， testing， development 和 other
       --print-json                                                       打印 JSON (default: false)
       --print-time                                                       打印时间 (default: false)
       --help, -h                                                         show help
    ```

- 导出（混合命令，多接口命令）

    ```shell
    $ go run main.go mix-export --help
    NAME:
       gitlab-go mix-export - 导出（混合命令，多接口命令）
    
    USAGE:
       gitlab-go mix-export command [command options]
    
    COMMANDS:
       all, a   导出所有（混合命令，多接口命令）
                已包含：
                1. git 仓库
                2. wiki 仓库
       help, h  Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value                                                   实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value                                                      your_access_token
       --owned                                                            当前用户明确拥有的项目。 (default: false)
       --export-folder value                                              导出文件夹
       --skip-project-path value [ --skip-project-path value ]            跳过项目路径
       --skip-project-wiki-path value [ --skip-project-wiki-path value ]  跳过项目wiki路径
       --auto-skip-exist-folder                                           自动跳过已存在的文件夹 (default: false)
       --allow-failure                                                    允许失败 (default: false)
       --help, -h                                                         show help
    ```

- 转移（混合命令，多接口命令）

    ```shell
    $ go run main.go mix-transfer --help
    NAME:
       gitlab-go mix-transfer - 转移（混合命令，多接口命令）
    
    USAGE:
       gitlab-go mix-transfer command [command options]
    
    COMMANDS:
       all, a   将一个命令空间的项目转移到新的命名空间
       help, h  Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value                                         实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value                                            your_access_token
       --owned                                                  当前用户明确拥有的项目。 (default: false)
       --namespace-source value                                 源命名空间。如：用户名、群组名
       --namespace-target value                                 目标命名空间。如：用户名、群组名
       --skip-project-path value [ --skip-project-path value ]  跳过项目路径
       --help, -h                                               show help
    ```

- 取消归档（混合命令，多接口命令）

    ```shell
    $ go run main.go mix-unarchive --help
    NAME:
       gitlab-go mix-unarchive - 取消归档（混合命令，多接口命令）
    
    USAGE:
       gitlab-go mix-unarchive command [command options]
    
    COMMANDS:
       all      取消归档所有项目（混合命令，多接口命令）
       help, h  Shows a list of commands or help for one command
    
    OPTIONS:
       --base-url value  实例地址，例如：https://gitlab.xuxiaowei.com.cn/api/v4 (default: "https://gitlab.com/api/v4") [%CI_API_V4_URL%]
       --token value     your_access_token
       --owned           当前用户明确拥有的项目。 (default: false)
       --help, -h        show help
    ```

### test

```shell
go test ./... -v
```

### build

```shell
go build
# GOOS=：设置构建的目标操作系统（darwin | freebsd | linux | windows）
# GOARCH=：设置构建的目标操作系统（386 | amd64 | arm | arm64）
# -v：打印编译过程中的详细信息
# -ldflags：设置在编译时传递给链接器的参数
# -ldflags "-s -w -buildid="
#                           -s: 删除符号表信息，减小可执行文件的大小。
#                           -w: 删除调试信息，使可执行文件在运行时不会打印调试信息。
#                           -buildid=: 删除构建ID，使可执行文件在运行时不会打印构建ID。
# -trimpath：去掉所有包含 go path 的路径
# -o：指定构建后输出的文件名
```

- Windows
    - amd64
        ```shell
        go build -o buildinfo/buildinfo.exe buildinfo/buildinfo.go
        GOOS=windows GOARCH=amd64 go build -v -ldflags "-s -w -buildid= -X main.BuildDate=$(buildinfo/buildinfo.exe now) -X main.Compiler= -X main.GitCommitBranch=$(buildinfo/buildinfo.exe commitBranch) -X main.GitCommitSha=$(buildinfo/buildinfo.exe commitSha) -X main.GitCommitShortSha=$(buildinfo/buildinfo.exe commitShortSha) -X main.GitCommitTag=$(buildinfo/buildinfo.exe commitTag) -X main.GitCommitTimestamp=$(buildinfo/buildinfo.exe commitTimestamp) -X main.GitTreeState=$(buildinfo/buildinfo.exe git-tree-state) -X main.GitVersion=$(buildinfo/buildinfo.exe commitTag) -X main.GoVersion=$(buildinfo/buildinfo.exe goShortVersion) -X main.Major= -X main.Minor= -X main.Revision= -X main.Platform=windows/amd64 -X main.CiPipelineId= -X main.CiJobId=" -trimpath -o gitlab-go-windows-amd64.exe .
        ```
    - arm64
        ```shell
        go build -o buildinfo/buildinfo.exe buildinfo/buildinfo.go
        GOOS=windows GOARCH=arm64 go build -v -ldflags "-s -w -buildid= -X main.BuildDate=$(buildinfo/buildinfo.exe now) -X main.Compiler= -X main.GitCommitBranch=$(buildinfo/buildinfo.exe commitBranch) -X main.GitCommitSha=$(buildinfo/buildinfo.exe commitSha) -X main.GitCommitShortSha=$(buildinfo/buildinfo.exe commitShortSha) -X main.GitCommitTag=$(buildinfo/buildinfo.exe commitTag) -X main.GitCommitTimestamp=$(buildinfo/buildinfo.exe commitTimestamp) -X main.GitTreeState=$(buildinfo/buildinfo.exe git-tree-state) -X main.GitVersion=$(buildinfo/buildinfo.exe commitTag) -X main.GoVersion=$(buildinfo/buildinfo.exe goShortVersion) -X main.Major= -X main.Minor= -X main.Revision= -X main.Platform=windows/arm64 -X main.CiPipelineId= -X main.CiJobId=" -trimpath -o gitlab-go-windows-arm64.exe .
        ```

- Linux
    - amd64
        ```shell
        go build -o buildinfo/buildinfo buildinfo/buildinfo.go
        GOOS=linux GOARCH=amd64 go build -v -ldflags "-s -w -buildid= -X main.BuildDate=$(buildinfo/buildinfo now) -X main.Compiler= -X main.GitCommitBranch=$(buildinfo/buildinfo commitBranch) -X main.GitCommitSha=$(buildinfo/buildinfo commitSha) -X main.GitCommitShortSha=$(buildinfo/buildinfo commitShortSha) -X main.GitCommitTag=$(buildinfo/buildinfo commitTag) -X main.GitCommitTimestamp=$(buildinfo/buildinfo commitTimestamp) -X main.GitTreeState=$(buildinfo/buildinfo git-tree-state) -X main.GitVersion=$(buildinfo/buildinfo commitTag) -X main.GoVersion=$(buildinfo/buildinfo goShortVersion) -X main.Major= -X main.Minor= -X main.Revision= -X main.Platform=linux/amd64 -X main.CiPipelineId= -X main.CiJobId=" -trimpath -o gitlab-go-linux-amd64 .
        ```
    - arm64
        ```shell
        go build -o buildinfo/buildinfo buildinfo/buildinfo.go
        GOOS=linux GOARCH=arm64 go build -v -ldflags "-s -w -buildid= -X main.BuildDate=$(buildinfo/buildinfo now) -X main.Compiler= -X main.GitCommitBranch=$(buildinfo/buildinfo commitBranch) -X main.GitCommitSha=$(buildinfo/buildinfo commitSha) -X main.GitCommitShortSha=$(buildinfo/buildinfo commitShortSha) -X main.GitCommitTag=$(buildinfo/buildinfo commitTag) -X main.GitCommitTimestamp=$(buildinfo/buildinfo commitTimestamp) -X main.GitTreeState=$(buildinfo/buildinfo git-tree-state) -X main.GitVersion=$(buildinfo/buildinfo commitTag) -X main.GoVersion=$(buildinfo/buildinfo goShortVersion) -X main.Major= -X main.Minor= -X main.Revision= -X main.Platform=linux/arm64 -X main.CiPipelineId= -X main.CiJobId=" -trimpath -o gitlab-go-linux-arm64 .
        ```

- LoongArch
    - 64-bit
        ```shell
        go build -o buildinfo/buildinfo buildinfo/buildinfo.go
        GOOS=linux GOARCH=loong64 go build -v -ldflags "-s -w -buildid= -X main.BuildDate=$(buildinfo/buildinfo now) -X main.Compiler= -X main.GitCommitBranch=$(buildinfo/buildinfo commitBranch) -X main.GitCommitSha=$(buildinfo/buildinfo commitSha) -X main.GitCommitShortSha=$(buildinfo/buildinfo commitShortSha) -X main.GitCommitTag=$(buildinfo/buildinfo commitTag) -X main.GitCommitTimestamp=$(buildinfo/buildinfo commitTimestamp) -X main.GitTreeState=$(buildinfo/buildinfo git-tree-state) -X main.GitVersion=$(buildinfo/buildinfo commitTag) -X main.GoVersion=$(buildinfo/buildinfo goShortVersion) -X main.Major= -X main.Minor= -X main.Revision= -X main.Platform=darwin/amd64 -X main.CiPipelineId= -X main.CiJobId=" -trimpath -o gitlab-go-loong64 .
        ```

- Darwin
    - amd64
        ```shell
        go build -o buildinfo/buildinfo buildinfo/buildinfo.go
        GOOS=darwin GOARCH=amd64 go build -v -ldflags "-s -w -buildid= -X main.BuildDate=$(buildinfo/buildinfo now) -X main.Compiler= -X main.GitCommitBranch=$(buildinfo/buildinfo commitBranch) -X main.GitCommitSha=$(buildinfo/buildinfo commitSha) -X main.GitCommitShortSha=$(buildinfo/buildinfo commitShortSha) -X main.GitCommitTag=$(buildinfo/buildinfo commitTag) -X main.GitCommitTimestamp=$(buildinfo/buildinfo commitTimestamp) -X main.GitTreeState=$(buildinfo/buildinfo git-tree-state) -X main.GitVersion=$(buildinfo/buildinfo commitTag) -X main.GoVersion=$(buildinfo/buildinfo goShortVersion) -X main.Major= -X main.Minor= -X main.Revision= -X main.Platform=darwin/amd64 -X main.CiPipelineId= -X main.CiJobId=" -trimpath -o gitlab-go-darwin-amd64 .
        ```
    - arm64
        ```shell
        go build -o buildinfo/buildinfo buildinfo/buildinfo.go
        GOOS=darwin GOARCH=arm64 go build -v -ldflags "-s -w -buildid= -X main.BuildDate=$(buildinfo/buildinfo now) -X main.Compiler= -X main.GitCommitBranch=$(buildinfo/buildinfo commitBranch) -X main.GitCommitSha=$(buildinfo/buildinfo commitSha) -X main.GitCommitShortSha=$(buildinfo/buildinfo commitShortSha) -X main.GitCommitTag=$(buildinfo/buildinfo commitTag) -X main.GitCommitTimestamp=$(buildinfo/buildinfo commitTimestamp) -X main.GitTreeState=$(buildinfo/buildinfo git-tree-state) -X main.GitVersion=$(buildinfo/buildinfo commitTag) -X main.GoVersion=$(buildinfo/buildinfo goShortVersion) -X main.Major= -X main.Minor= -X main.Revision= -X main.Platform=darwin/arm64 -X main.CiPipelineId= -X main.CiJobId=" -trimpath -o gitlab-go-darwin-arm64 .
        ```

## 鸣谢

1. 感谢 [![jetbrains](./static/jb_beam.svg)](https://www.jetbrains.com/)
   提供开发工具 [![IDEA](./static/IntelliJ_IDEA_icon.svg)](https://www.jetbrains.com/idea) 的免费授权

## Stargazers over time

[![Stargazers over time](https://starchart.cc/xuxiaowei-com-cn/gitlab-go.svg)](https://starchart.cc/xuxiaowei-com-cn/gitlab-go)
