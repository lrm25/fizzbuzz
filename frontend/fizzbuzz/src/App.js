import './App.css';
import React, { useState } from 'react';

function App () {
  const [count, setCount] = useState(0);
  const [message, setMessage] = useState('');

  function postToBackend() {
    fetch(process.env.REACT_APP_BACKEND_URL + '/fizzbuzz',
      {method: 'POST',
       body: JSON.stringify({
        "count": count + 1
       }),
       headers: {'Content-type': 'application/json'}})
      .then(async response => {
        const responseJson = await response.json();
        if (!response.ok) {
          const error = responseJson || response.status;
          alert(error);
          return;
        }
        setMessage(responseJson["message"]);
        setCount(count + 1);
      })
      .catch(error => alert(error));
  }

  return (
    <div className="app">
      <div className="text count-desc">Your Count</div>
      <div className="text count-val">{count}</div>
      <button onClick={postToBackend} className="button">Push me!</button>
      <h1 className="fizz-buzz">{message}</h1>
    </div>
  );
}

export default App;
