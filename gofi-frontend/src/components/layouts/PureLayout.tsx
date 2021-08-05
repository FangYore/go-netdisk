import React from 'react'
import Footer from '../Footer'

const PureLayout: React.FC = (props) => {
    return (
        <div className="h-full flex flex-col items-center justify-center space-y-6 p-6">
            <div className="flex-grow flex flex-col items-center justify-center w-full h-full">{props.children}</div>
            <Footer />
        </div>
    )
}

export default PureLayout
