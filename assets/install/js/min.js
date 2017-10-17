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
}

$(document).ready(function(){
    init_smartwizard();
    //bind_widget_events();
});