#include<iostream>
#include<thread>
#include<vector>

int findLowerBound(std::vector<int>& arr, int target){
    int left = 0;
    int right = arr.size() - 1;
    while (left <= right){
        int mid = left + (right - left)/2;
        if(arr[mid] < target){
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return left;
}

int main(){ 
    std::cout << " == Given a sorted array, find the range of elements with value = target" << std::endl;
	std::vector<int> arr = {1,2,3,4,4,4,4,5,6,7,7,8,9};
	int target = 8;
    int leftBound, rightBound;
    //
    std::thread t1([&](){
        leftBound = findLowerBound(arr,target);
    });
    std::thread t2([&](){
        rightBound = findLowerBound(arr,target+1) - 1;
    });
    t1.join();
    t2.join();
    if(leftBound > rightBound){
        std::cout << "[-1,-1]" << std::endl;
    } else {
        std::cout << "[" << leftBound << "," << rightBound << "]" << std::endl;
    }
    return 0;
}