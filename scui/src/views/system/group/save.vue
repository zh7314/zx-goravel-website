<template>
	<el-dialog :title="title" v-model="visible" :width="500" destroy-on-close @closed="$emit('closed')">
		<el-form :model="form" :rules="rules" ref="dialogForm" label-width="100px">
			<el-form-item label="上级部门" prop="parent_id">
				<el-cascader v-model="form.parent_id" :options="groups" :props="groupsProps" :show-all-levels="false"
							 clearable style="width: 100%;"></el-cascader>
			</el-form-item>
			<el-form-item label="部门名称" prop="name">
				<el-input v-model="form.name" placeholder="请输入部门名称" clearable></el-input>
			</el-form-item>
			<el-form-item label="排序" prop="sort">
				<el-input-number v-model="form.sort" controls-position="right" :min="1"
								 :max="255" style="width: 30%;"></el-input-number>
			</el-form-item>
		</el-form>
		<template #footer>
			<el-button @click="visible=false">取 消</el-button>
			<el-button type="primary" :loading="isSaveing" @click="submit()">保 存</el-button>
		</template>
	</el-dialog>
</template>

<script>
export default {
	emits: ['success', 'closed'],
	data() {
		return {
			title: '编辑',
			visible: false,
			isSaveing: false,
			//表单数据
			form: {
				id: "",
				parent_id: "",
				name: "",
				sort: 255
			},
			//验证规则
			rules: {
				sort: [
					{required: true, message: '请输入排序', trigger: 'change'}
				],
				name: [
					{required: true, message: '请输入部门名称'}
				]
			},
			topOptions: {
				id: 0,
				value: 0,
				parent_id: 0,
				label: "顶级分类",
				name: "顶级分类",
				children: []
			},
			//所需数据选项
			groups: [],
			groupsProps: {
				value: "id",
				emitPath: false,
				checkStrictly: true,
				label: "name"
			}
		}
	},
	mounted() {
		this.getGroup()
	},
	methods: {
		//显示
		open() {
			this.visible = true;
			return this
		},
		//加载树数据
		async getGroup() {
			var res = await this.$API.system.adminGroup.getGroupTree.get();
			this.groups = res.data;
			this.groups.unshift(this.topOptions)
		},
		//表单提交方法
		submit() {
			this.$refs.dialogForm.validate(async (valid) => {
				if (valid) {
					this.isSaveing = true;
					var res = await this.$API.system.adminGroup.save.post(this.form);
					this.isSaveing = false;
					if (res.code == 200) {
						this.$emit('success', this.form)
						this.visible = false;
						this.$message.success("操作成功")
					} else {
						this.$alert(res.msg, "提示", {type: 'error'})
					}
				}
			})
		},
		//表单注入数据
		setData(data) {
			//可以和上面一样单个注入，也可以像下面一样直接合并进去
			Object.assign(this.form, data)
			this.form.permission_ids = null
		}
	}
}
</script>

<style>
</style>
