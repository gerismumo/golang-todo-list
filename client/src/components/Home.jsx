import React, { useState } from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {faTrashCan} from '@fortawesome/free-solid-svg-icons';
import {faPenToSquare } from '@fortawesome/free-regular-svg-icons';
import axios from 'axios';
import { Db_url } from './utils';


const Home = () => {
    const[openEdit, setOpenEdit] = useState(false);
    
    const[todoName, setTodoName] = useState('');
    const[searchTodo, setSearchTodo] = useState('');

    const handleAddTodo = async(e) => {
        e.preventDefault();

        try {
            const data = {
                todoName,
            }

            const response = await axios.post(`${Db_url}/api/addTodo`,data);
            console.log(response.data);

        }catch(error) {
            console.log(error.message);
        }

    }

  return (
    <div className="flex flex-row justify-center py-[40px]">
        <div className="flex flex-col justify-center items-center ">
            <div className="">
                <h1 className='text-[35px] font-[700]'>Golang Todo</h1>
            </div>
            <div className="flex flex-row justify-between items-center  gap-[100px] mt-[50px]">
                    <form action="" className='flex flex-row gap-[20px] items-center '>
                        <input type="text" 
                        placeholder="Search Todo" 
                        className=' py-[5px] px-[4px] border-[1px] border-[#068FFF] rounded-[6px] outline-none '
                        />
                        {/* <button type="Search"
                        className='bg-[#4E4FEB] rounded-[6px] text-[#ffffff] py-[5px] px-[20px] '
                        >Search</button> */}
                    </form>
                    <form onSubmit={(e) => handleAddTodo(e)} className='flex flex-row gap-[20px] items-center '>
                        <input type="text" 
                        name='todo'
                        value={todoName}
                        onChange={(e) => setTodoName(e.target.value)}
                        placeholder="Add Todo" 
                        className='py-[5px] px-[4px] border-[1px] border-[#068FFF] rounded-[6px] outline-none '
                        />
                        <button type="submit"
                        className='bg-[#19A7CE] rounded-[6px] text-[#ffffff] py-[5px] px-[20px] '
                        >Add</button>
                    </form>
            </div>
            <div className="mt-[70px]">
                <table className='border-collapse  '>
                    <thead>
                        <th className='border border-[#ddd] px-[30px] py-[10px]'>Id</th>
                        <th className='border border-[#ddd] px-[30px] py-[10px]'>Name</th>
                        <th className='border border-[#ddd] px-[30px] py-[10px]'>Created At</th>
                    </thead>
                    <tbody>
                        <tr>
                            <td className='border border-[#ddd] px-[30px] py-[10px]'>1</td>
                            <td className='border border-[#ddd] px-[30px] py-[10px]'>Gera</td>
                            <td className='border border-[#ddd] px-[30px] py-[10px]'>789</td>
                            <td className='border border-[#ddd] px-[30px] py-[10px]'><FontAwesomeIcon color='red' icon={faTrashCan} /></td>
                            <td className='border border-[#ddd] px-[30px] py-[10px]'><FontAwesomeIcon color='#068FFF' icon={faPenToSquare} /></td>
                        </tr>
                        <tr>
                            <td className='border border-[#ddd] px-[30px] py-[10px]'>2</td>
                            <td className='border border-[#ddd] px-[30px] py-[10px]'>Gera</td>
                            <td className='border border-[#ddd] px-[30px] py-[10px]'>789</td>
                            <td className='border border-[#ddd] px-[30px] py-[10px]'><FontAwesomeIcon color='red' icon={faTrashCan} /></td>
                            <td className='border border-[#ddd] px-[30px] py-[10px]'><FontAwesomeIcon color='#068FFF' icon={faPenToSquare} /></td>
                        </tr>
                        <tr>
                            <td className='border border-[#ddd] px-[30px] py-[10px]'>3</td>
                            <td className='border border-[#ddd] px-[30px] py-[10px]'>Gera</td>
                            <td className='border border-[#ddd] px-[30px] py-[10px]'>789</td>
                            <td className='border border-[#ddd] px-[30px] py-[10px]'><FontAwesomeIcon color='red' icon={faTrashCan} /></td>
                            <td className='border border-[#ddd] px-[30px] py-[10px]'><FontAwesomeIcon color='#068FFF' icon={faPenToSquare} /></td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
        
    </div>
  )
}

export default Home