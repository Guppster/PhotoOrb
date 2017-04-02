const apiUrl = 'http://ec2-54-175-181-19.compute-1.amazonaws.com';
const user = "nick"

getUserImages = $.ajax(
    {
        type: "GET",
        url: apiUrl + "/images/" + user
    }
);

getUserImages.done(function (response) {
    console.log(response);
});
getUserImages.fail(function (response) {
    console.log(response);
});
