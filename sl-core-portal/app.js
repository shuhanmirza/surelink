const app = Vue.createApp({
    data() {
        return {
            firstName:'john',
            lastName:'doe',
            email: 'john.doe@gmail.com',
            gender: 'male',
            picture:'https://randomuser.me/api/portraits/men/10.jpg'
        }
    },
    methods: {

        async getUser() {
            let res = await fetch('https://randomuser.me/api')
            let {results} = await res.json()
            console.log(results)

            let user = results[0]
            console.log(user)

            this.firstName = user.name.first
            this.lastName = user.name.last
            this.gender = user.gender
            this.email = user.email
            this.picture = user.picture.large
        }
    },
})

app.mount('#app');