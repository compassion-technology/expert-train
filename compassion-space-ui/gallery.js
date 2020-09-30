$(document).ready(function(){
    $.getJSON( 'http://ec2co-ecsel-6n3bncmeboii-471081296.us-east-2.elb.amazonaws.com/user/01234567/messages',function( data ) {
        $.each( data, function( val ) {
            if (data[val].message.content_url != null) {
                $(".row").append("<div class='col-lg-4 col-md-12 mb-4'><div class='modal fade' id='modal1' tabindex='-1' role='dialog' aria-labelledby='myModalLabel' aria-hidden='true'><div class='modal-dialog modal-lg' role='document'><div class='modal-content'><div class='modal-body mb-0 p-0'><div class='embed-responsive embed-responsive-16by9 z-depth-1-half'><iframe class='embed-responsive-item' src='"+data[val].message.content_url+"'></iframe></div></div><div class='modal-footer justify-content-center'><button type='button' class='btn btn-outline-primary btn-rounded btn-md ml-4' data-dismiss='modal'>Close</button></div></div></div></div><a><img class='img-fluid z-depth-1' src='"+data[val].message.content_url+"' style= 'width: 100%;' alt='video' data-toggle='modal' data-target='#modal1'></a></div>");
            }
         });
});
});