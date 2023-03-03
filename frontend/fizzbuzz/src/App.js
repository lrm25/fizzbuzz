import './App.css';
import {useState} from 'react';

function App() {

  const [count, setCount] = useState(0);
  const [message, setMessage] = useState("");

  function postToBackend() {
    fetch(process.env.REACT_APP_BACKEND_URL + '/fizzbuzz',
      {method: 'POST',
       body: JSON.stringify({
        "count": count + 1
       }),
       headers: {'Content-type': 'application/json'}})
      .then(async response => {
        const responseJson = await response.json()
        console.log(responseJson)
        if (!response.ok) {
          const error = responseJson || response.status;
          alert(error);
          return;
        }
        setMessage(responseJson["message"])
        setCount(count + 1)
      })
      .catch(error => alert(error));
  }

  return (
    <div className="App">
      <div className="Text">Your Count</div>
      <div className="Text">{count}</div>
      <button onClick={postToBackend} className="Button">Push me!</button>
      <h1 className="FizzBuzz">{message}</h1>
    </div>
  );
}

export default App;
