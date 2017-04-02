const apiUrl = 'http://ec2-54-175-181-19.compute-1.amazonaws.com';
const user = "nick";

getUserImages = $.ajax(
  {
    type: "GET",
    url: apiUrl + "/images/" + user
  }
);


getUserImages.done(function (response) {
  console.log(response);
  for (var i = 1; i <= response.length; i++) {
    $.ajax(
      {
        type: "GET",
        url: apiUrl + "/images/" + user + "/" + response[i]
      }
    ).done(function (response) {
      var toAdd =
        ' <figure class="gallery-item  text-center homeFigure" style="width: auto;padding: 0px;margin: 0px"> ' +
        ' <img src="' + apiUrl + '/images/nick/' + 1 + '/11.jpg" height="auto" width="100%"' +
        ' class="reel center-block"' +
        ' id="image2"' +
        ' data-images="http://ec2-54-175-181-19.compute-1.amazonaws.com/images/nick/' + 1 + '/##.jpg|11..' + 30 + '"' +
        ' data-loops="false"' +
        ' class="image-responsive">' +
        ' <br>' +
        ' </figure>';

      console.log(response)

      // document.getElementById("previewSection").innerText(toAdd);

      // $("#previewSection").append(toAdd);
    })
  }


});

getUserImages.fail(function (response) {
  console.log(response);
});
