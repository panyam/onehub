
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
    return this.setLoggedInUser(id, {
      name: name,
    })
  }

  logout() {
    localStorage.setItem("loggedInUser", JSON.stringify({}))
  }

  setLoggedInUser(id: string, profile: any) {
    profile = profile || {}
    id = id.trim()
    profile.id = id
    localStorage.setItem("loggedInUser", JSON.stringify(profile))
    return profile
  }

  get loggedInUser() {
    if (typeof (localStorage) === "undefined") {
      return {'id': null}
    }
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
