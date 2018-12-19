function change() {
    var height01 = $(window).height();
    $(".top").css('height', height01 - 35+"px");
}

!function () {
    laydate.skin('dahong');
    laydate({ elem: '#CompanyInTime' });
    laydate({ elem: '#CompanyOutTime' });
}();

$(document).ready(function () {
    // init();
    $("#form_company_add_save").click(saveOrUpdateCompany);
});

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
                }
                else {
                    alert("获取失败！");
                }
            },
            error: function (err) {
            }
        })
    }
}
function saveOrUpdateCompany() {
    if($("#CompanyId").val()) {
        updateCompany();
    } else {
        saveCompany();
    }
}

function saveCompany() {
    $.ajax({
        type: "POST",
        url: "/company/add",
        data: $("#form_company").serialize(),
        dataType: "json",
        success: function (result) {
            if (result.status == 1) {
                alert("保存成功！！！");
                parent.getCompanyTableData();
                TaskCancel();
            } else {
                alert("保存失败！！！")
            }
        }
    })
}

function updateCompany() {

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