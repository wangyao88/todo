function change() {
    var height01 = $(window).height();
    $(".top").css('height', height01 - 35+"px");
}

var PROJECTID, TASKTYPE, TASKPHASE,  PERSONID
!function () {
    laydate.skin('danlan');//切换皮肤，请查看skins下面皮肤库
    laydate({ elem: '#demo' });//绑定元素
    laydate({ elem: '#demo1' });
    laydate({ elem: '#demo2' });//绑定元素
}();
function init() {
    if (parent.getCurrentID() != "") {
        $.ajax({
            url: '../TaskRecord/SearchTaskById?taskId=' + parent.getCurrentID(),
            type: 'POST',
            dataType: 'json',
            success: function (result) {
                var data = result.data;
                if (data) {
                    PROJECTID = data.PROJECTID
                    PERSONID =data.PERSONID
                    TASKPHASE = data.TASKPHASE;
                    TASKTYPE = data.TASKTYPE;
                    $("#tit").val(data.PROJECTNAME);
                    $("#person").val(data.TASKPHASENAME);
                    $("#part").val(data.STATE);

                    $("#demo").val(getFormatTime(data.STARTTIME.substring(6,19)));


                    $("#NOTE").val(data.NOTE);

                    getPart();

                }
                else {
                    alert("获取失败！");
                }
            },
            error: function (err) {
            }
        })

    }
    else {
        getPart();
    }
}
function NoticeSave() {
    $.ajax({
        type: "POST",
        url: "../TaskRecord/UpdateTask",
        data: {
            TASKID: parent.getCurrentID(),
            PROJECTID: PROJECTID,
            STATE: $("#tit").val(),
            STARTTIME: $("#demo").val(),

            PERSONID: PERSONID,
            NOTE: $("#NOTE").val(),
            TASKTYPE: $("#part").val(),
            TASKPHASE: $("#person").val(),

        },
        dataType: "json",
        success: function (result) {
            if (result.data) {
                alert("保存成功！！！");
                parent.getNoticeTableData();
                TaskCancel();
            } else {
                alert("保存失败！！！")
            }
        }
    })
}
function TaskCancel() {
    var index = parent.layer.getFrameIndex(window.name)
    parent.layer.close(index);
}

//时间格式化函数
function getFormatTime(time) {
    var time = new Date(parseInt(time));
    var y = time.getFullYear();
    var m = time.getMonth() + 1;
    var d = time.getDate();
    var h = time.getHours();
    var mm = time.getMinutes();
    var s = time.getSeconds();
    return y + '-' + add0(m) + '-' + add0(d) + ' ' + add0(h) + ':' + add0(mm) + ':' + add0(s);
}
function add0(m) { return m < 10 ? '0' + m : m }