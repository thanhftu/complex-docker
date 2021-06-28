import {useState, useEffect} from 'react';
import axios from 'axios';
import Fib from './Fib'

const Fibs = ()=>{
    // const [seenIndex, setSeenIndex]=useState([])
    const [values, setValues]=useState([])
    const [index, setIndex]=useState('')
    const [calculatedIndex,setCalculatedIndex]=useState([])
  useEffect(()=>{
    const getFibs=async ()=>{
      const fibsFromServer=await fetchValues()
      setValues(fibsFromServer)
    }
    getFibs()
  },[])
    const fetchValues = async ()=>{
        const getdata = await axios.get('http://localhost:8081/api/allfib')
        const fibs = getdata.data
        return fibs
    }
  console.log(values)
    console.log(values)

    const handleSubmit = async (e) => {
        e.preventDefault();
        setCalculatedIndex([...calculatedIndex,index])
        const res=await axios.post('http://localhost:8081/api/values', {
          index: index
        });
        console.log("res",res.data)
        setValues([...values,res.data])
        setIndex('');
      };
      console.log("cal is",calculatedIndex)
      console.log("val is",values)
    
      const onDelete = async (id) => {
        const options = {
          headers: {
              'Content-Type': 'application/json',
          }
        };
        const newVals = values.filter(val=>val.ID!==id)
        setValues(newVals);
        await axios.delete(`http://localhost:8081/api/values/${id}`, {
          id: id
        },options);
        console.log(id)
        
      };
   
    // const renderSeenIndexes() {
    //     return seenIndexes.map(({ number }) => number).join(', ');
    //   }
    
    // const renderValues=() {
    //     const entries = [];
    
    //     for (let key in values) {
    //       entries.push(
    //         <div key={key}>
    //           For index {key} I calculated {values[key]}
    //         </div>
    //       );
    //     }
    
    //     return entries;
    //   }
    return (
      <div>
        <form onSubmit={handleSubmit}>
          <label>Enter your index:</label>
          <input
            value={index}
            onChange={(e) => setIndex(e.target.value)}
          />
          <button>Submit</button>
        </form>

        
        <h3>Values:</h3>
        {values.map((fib, index)=><Fib key={index} fib={fib} onDelete={onDelete} text='Delete'/>)}
        <h3>calculated Index:</h3>
        {calculatedIndex.map((cal)=><h4>{cal}</h4>)}
        
      </div>
        // <div>
        //     <form>
        //         <label>
        //             Enter your index
        //         </label>
        //         <input 
        //             // value={index}
        //             // onChange={event => setIndex({ index: event.target.value })}
        //         />
        //         <button>Submit</button>
        //     </form>
        //     <h3>Indexes I have seen:</h3>
        //     {/* {renderSeenIndexes()} */}

        // <h3>Calculated Values:</h3>
        // {renderValues()}
        // </div>
    )
}

export default Fibs;