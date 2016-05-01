
$( document ).ready(function() {

  base = "http://localhost:8080/post";

  $("#post").click(function(){
    message = $("#message").val();
    if(message.length === 0) {
      alert("Missing status");
      return;
    }
      var status = encodeURIComponent(message);

      $.ajax({
        url: base +"?status="+message,
        data: "",
        success: function(m){
          console.log(JSON.stringify(m));
        }
      });
  });

});
