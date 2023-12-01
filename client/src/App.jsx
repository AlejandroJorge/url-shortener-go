import { useReducer, useRef, useState } from "react";

export default function App(){
  const [requestState, setRequestState] = useState("idle")
  const [returnedURL, setReturnedURL] = useState("")
  const originalURL = useRef(null)
  
  async function makeRequest(){
    console.log("making request to server", createRequest(), import.meta.env.VITE_API_URL)
    fetch(`${import.meta.env.VITE_API_URL}`,createRequest())
      .then(response => response.json())
      .then(data => {
        setReturnedURL(data.shortenedURL)
        setRequestState("success")
      })
      .catch(error => {
        setRequestState("error")
        console.log(error)
      })
  }

  async function reset(){
    setRequestState("idle")
  }

  function createRequest(){
    return {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        originalURL: originalURL.current.value 
      })
    }
  }

  return <>
    <main style={{display:"flex", flexDirection:"column"}}>
      <h1>URL Shortener App</h1>
      {requestState === "idle" && <Form originalURL={originalURL} handleClick={makeRequest}></Form>}
      {requestState === "success" && <Result returnedURL={returnedURL} handleClick={reset}></Result>}
      {requestState === "error" && <Error handleClick={reset}></Error>}
   
     </main>
  </>
}

function Form({originalURL, handleClick}){
  return <>
    <input ref={originalURL} type="text" placeholder="https://" />
    <button onClick={handleClick}>Shorten</button>
  </>
}

function Result({returnedURL, handleClick}){
  return <>
    <p>Your shortened URL: {returnedURL}</p>
    <button onClick={handleClick}>Try again</button>
  </>
}

function Error({handleClick}){
  return <>
    <p>There was an error</p>
    <button onClick={handleClick}>Try again</button>
  </>
}