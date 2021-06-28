import PropTypes from 'prop-types'

const Button = ({fib, onClick, text})=>{
    return (
        <button
        className='btn'
            onClick={()=>onClick(fib.ID)}
        >
            {text}
        </button>
    )
}

Button.protypes={
    onClick:PropTypes.func.isRequired,
    text:PropTypes.string,
}
export default Button