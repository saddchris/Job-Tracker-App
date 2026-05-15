import { useEffect, useState } from "react"

function App() {
  const [message, setMessage] = useState("Loading...")

  useEffect(() => {
    fetch("http://localhost:8080/api")
      .then(res => res.json())
      .then(data => setMessage(data.message))
      .catch(() => setMessage("Error connecting to backend"))
  }, [])

  return (
    <div style={{ textAlign: "center", marginTop: "50px" }}>
      <h1>Full Stack App</h1>
      <p>{message}</p>
    </div>
  )
}

export default App