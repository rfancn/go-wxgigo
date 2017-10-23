function validate_step(stepNumber){
    var form = $('#form' + stepNumber);
    return form.parsley().validate();
}

function init_step_server(){
    if($.trim(server.host).length === 0){
        server.host = wxmp.host;
    }
}

var progress = {
    //flag to indicate whether it need to be destroyed or not
    //if it is true, then it will not start a new setTimeout func
    todestory: false,
    tid: -1,
    interval: 1000,
    init: function(){
        waitingDialog.show('Progressing...', {dialogSize: 'sm'});
        this.todestory = false;
        this.start();
    },
    // kicks off the setTimeout
    start: function(){
        this.tid = setTimeout(
            $.proxy(this.fetch, this), // ensures 'this' is the poller obj inside getData, not the window object
            this.interval
        );
    },
    // get AJAX data + respond to it
    fetch: function(){
       var self = this;
       $.ajax({
           url: '/install/progress',
           success: function(data){
                console.log("fetch success:" + $.now());
                waitingDialog.progress(data);
           },
           complete: function(){
               //if this progress don't need to be destroyed,
               // then continue start a timer
                if (!self.todestory) {
                    self.start();
                }
           },
       });
    },
    // stop timeout
    stop: function(){
        console.log("stop progress");
        //mark no new setTimeout()
        this.todestory = true;
        clearTimeout(this.tid);
    },
    destroy: function(){
        console.log("destroy progress");
        //mark no new setTimeout()
        this.todestory = true;
        clearTimeout(this.tid);
        waitingDialog.hide();
    },
    finish: function(text){
        console.log("finish progress");
        if(typeof text !== 'undefined' ){
            waitingDialog.message(text);
        }else{
            waitingDialog.message("Complete Successfully!");
        }
        waitingDialog.percentage(100);
    },

};

function getHttpsUrl(pathname) {
    var withSlash = true;
    if (pathname.substring(0, 1) == "/") {
        withSlash = false;
    }
    var httpsUrl = "https://" + window.location.hostname;
    if (withSlash) {
        httpsUrl += "/" + pathname;
    }else{
        httpsUrl += pathname;
    }

    return httpsUrl
}

function on_click_finish(){
    // validate all forms
    var errorSteps = [];
    // step index, only step1,2 and 3 need to be verified
    //steps = [0, 1, 2]
    steps = []
    for(var i=0;i<steps.length;i++){
        if(!validate_step(steps[i])){
            errorSteps.push(i+1);
        }
    }

    if(errorSteps.length > 0){
        //$('#smartwizard').smartWizard("stepState", errorSteps, "error");
        alertify.alert("You need correct the errors in step"+errorSteps);
        return false;
    }

    $.ajax({
            //url: getHttpsUrl("/install/save"),
            //crossDomain: true,
            url: "/install/save",
            dataType: 'json',
            contentType:"application/json; charset=utf-8",
            type: 'POST',
            data: JSON.stringify({
                'general': general.$data,
                'wxmp': wxmp.$data,
                'celery': celery.$data,
                'server': table_servers.rows().data().toArray(),
            }),
            beforeSend: function(xhr, settings) {
                progress.init();
            },
            success: function(data) {
                console.log("save success");
                console.log(data);
                progress.finish();
                //setTimeout(function(){
                //    progress.destroy();
                //    //window.location.replace(data);
                //}, 2000);
            },
            error: function(xhr, status){
                console.log("save error");
                progress.destroy();
            }
    });

    return true;
}

function init_smartwizard(){
    $('#smartwizard').smartWizard({
        keyNavigation:false, // Enable/Disable keyboard navigation(left and right keys are used if enabled)
        toolbarSettings: {
            toolbarExtraButtons: [
                $('<button id="btn-finish"></button>').text('Finish').addClass('btn btn-primary').on('click', function(){
                    on_click_finish();
                }),
            ],
        },
    });

    $('#btn-finish').hide();

    // Initialize the leaveStep event
    $("#smartwizard").on("leaveStep", function(e, anchorObject, stepNumber, stepDirection) {
        var isStepValidated = true;

        // validate html when forward to next step
        if(stepDirection === 'forward'){
            isStepValidated = validate_step(stepNumber);
        }

        if(isStepValidated == false){
            alertify.alert("You need correct the errors marked as red to continue!");
            return false;
        }

        return true;
    });

    // Initialize the showStep event
    $("#smartwizard").on("showStep", function(e, anchorObject, stepNumber, stepDirection) {
        // if we change to previous steps, then we need validate again
        switch(stepNumber){
            case 3:
                $('#btn-finish').show();
                init_step_server();
                // show finish button and set it's enabled status if validation is ok
                break;
        }
    });
}

$(document).ready(function(){
    init_smartwizard();
});

