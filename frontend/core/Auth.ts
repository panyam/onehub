
export default class Auth {
  ensureLoggedIn() {
    if (this.isLoggedIn) {
      return this.loggedInUser
    }
    let name = prompt("Enter your name")
    if (name == null || name.trim() == "") {
      return null
    }
    name = name.trim()
    let id = prompt("Enter your user id")
    if (id == null || id.trim() == "") {
      return null
    }
    id = id.trim()
    const user = {
      name: name,
      id: id,
    }
    localStorage.setItem("loggedInUser", JSON.stringify(user))
    return user
  }

  get loggedInUser() {
    let out = localStorage.getItem("loggedInUser")
    if (out == null) {
      const val = {
        id: null,
        name: "",
      }
      out = JSON.stringify(val)
      localStorage.setItem("loggedInUser", out)
    }
    return JSON.parse(out)
  }

  get isLoggedIn() {
    return this.loggedInUser.id != null
  }
}
