import React from 'react';
import './shortener.css';

class Shortener extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            long: null,
            short: null
        };
    }

    render() {
        return (
            <div className="container">
                <input></input>
                <button>Copy</button>
                <div className="links">
                    <p>Shortened Link</p>
                    <p>Longer Link this is long</p>
                </div>
                
            </div>
        )
    }
}

export default Shortener;