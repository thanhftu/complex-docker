import './App.css';
import Fibs from './Fibs'
import OtherPage from './OtherPage'
import {BrowserRouter as Router, Route, Link} from "react-router-dom";

function App() {
  return (
    <Router>
    <div className="App">
          <header>
            <h1 className="App-title">Calculate Fibonaci</h1>
            <Link to="/">Home</Link>
            <Link to="/otherpage">Other Page</Link>
          </header>
          <div>
            <Route exact path="/" component={Fibs} />
            <Route path="/otherpage" component={OtherPage} />
          </div>
      </div>

  </Router>
    );
}

export default App;
