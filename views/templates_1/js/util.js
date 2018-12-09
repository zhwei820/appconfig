/**
 * Created by zw on 2017/12/8.
 */
function getURLParameter(name) {
    return decodeURIComponent((new RegExp('[?|&]' + name + '=' + '([^&;]+?)(&|#|;|$)').exec(location.search) || [null, ''])[1].replace(/\+/g, '%20')) || null;
}

function getUrlParamValue(name) {
    return decodeURIComponent((new RegExp('[?|&]' + name + '=' + '([^&;]+?)(&|#|;|$)').exec(location.search) || [null, ''])[1].replace(/\+/g, '%20')) || null;
}

function diffUsingJS(_old, _new) {
    "use strict";
    var byId = function (id) {
        return document.getElementById(id);
    };
    var base = difflib.stringAsLines(_old),
        newtxt = difflib.stringAsLines(_new),
        sm = new difflib.SequenceMatcher(base, newtxt),
        opcodes = sm.get_opcodes(),
        diffoutputdiv = byId("diffoutput"),
        contextSize = 16;

    diffoutputdiv.innerHTML = "";
    contextSize = contextSize || null;

    diffoutputdiv.appendChild(diffview.buildView({
        baseTextLines: base,
        newTextLines: newtxt,
        opcodes: opcodes,
        baseTextName: "Base Text",
        newTextName: "New Text",
        contextSize: contextSize,
        viewType: 0
    }));
}


function display_diff(layer, $, _old, _new) {
    diffUsingJS(_old, _new)

    layer.open({
        type: 1,
        title: '修改详情',              // [可选]
        content: $('#diff'),          // 对话框中的内容部分
        area: ['1500px',],    // 对话框的大小 [可选]
        shadeClose: true,           // 为 true 时点击遮罩关闭 [可选]
    });
}


//写cookies
function setCookie(name, value) {
    var Days = 30;
    var exp = new Date();
    exp.setTime(exp.getTime() + Days * 24 * 60 * 60 * 1000);
    document.cookie = name + "=" + escape(value) + ";expires=" + exp.toGMTString();
}

//读取cookies
function getCookie(name) {
    var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");

    if (arr = document.cookie.match(reg))
        return unescape(arr[2]);
    else
        return null;
}

//删除cookies
function delCookie(name) {
    var exp = new Date();
    exp.setTime(exp.getTime() - 1);
    var cval = getCookie(name);
    if (cval != null)
        document.cookie = name + "=" + cval + ";expires=" + exp.toGMTString();
}