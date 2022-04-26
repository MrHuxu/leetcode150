struct Solution;

impl Solution {
    pub fn can_complete_circuit(gas: Vec<i32>, cost: Vec<i32>) -> i32 {
        let mut total_gas = 0;
        let mut total_cost = 0;

        for i in 0..gas.len() {
            total_gas += gas[i];
            total_cost += cost[i];
        }

        if total_cost > total_gas {
            return -1;
        }

        let mut ret = 0;

        let mut gas_sum = 0;
        let mut cost_sum = 0;
        for i in 0..gas.len() {
            gas_sum += gas[i];
            cost_sum += cost[i];
            if cost_sum > gas_sum {
                gas_sum = 0;
                cost_sum = 0;
                ret = i + 1;
            }
        }

        ret as i32
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::can_complete_circuit(vec![1, 2, 3, 4, 5], vec![3, 4, 5, 1, 2]),
        3
    );
    assert_eq!(
        Solution::can_complete_circuit(vec![2, 3, 4], vec![3, 4, 3]),
        -1
    );
}
