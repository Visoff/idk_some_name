<script>
    import req from '@lib/fetch'

    let url = "#"

    req("/bucket/myfile", {
        headers:{
            Authorization:"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoidXNlcjEyMyJ9.dC_hrWMOlI2Pg9M-Chgg4DIwZjJA7RzA0BeKhBpYlgE"
        }
    }).then(response => response.blob())
    .then(blob => {
        url = URL.createObjectURL(blob)
    })
    .catch(console.error)
</script>

<img src={url} alt="">