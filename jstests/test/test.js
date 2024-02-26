let apiUrl = `http://localhost:6700`
const { assert } = require('chai')
const rp = require('request-promise').defaults({ resolveWithFullResponse: true })
console.log("hello")
let curEmail 
let pw = "Password1"
let tok 
let allP
let authHeaders

function mkauth(token) {
	let x = {
		authorization: `Bearer ${token}`
	}
	authHeaders = x
}
describe('basic endpoints', async function () {
	it("should get /", async function () {
		let opts = { url: apiUrl }
		let res = await rp(opts)
		console.log(res.statusCode)
		assert.equal(res.statusCode, 200)
	})

	it('should register', async function () {
		let suffix = Date.now().toString().slice(-5)
		let email = `myemail${suffix}@gmail.com`
		curEmail = email
		let opts = {
			url: `${apiUrl}/register`,
			json: {
				email: email, password: pw
			}
		}
		let res = await rp(opts)
		console.log(res.body)
		assert.equal(res.statusCode, 200)
		assert.equal(res.body.success, true)
		assert.isString(res.body.token)
		tok = res.body.token
		console.log(tok)
		mkauth(tok)
	})
	it('should login', async function () {
		let opts = {
			url: `${apiUrl}/login`,
			json: {
				email: curEmail,
				password: pw,

			}
		}
		let res = await rp(opts)
		// console.log(res.body)
		assert.equal(res.body.success, true)
		assert.isString(res.body.token)
		tok = res.body.token
		mkauth(tok)

	})
	// return
	
	let todelete = null
})
